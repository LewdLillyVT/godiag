package main

import (
	"GoDiag/modules"
	"GoDiag/rpc"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	currentVersion = "1.0.7"
	updateCheckURL = "https://raw.githubusercontent.com/LewdLillyVT/godiag/refs/heads/main/version.json"
)

type VersionInfo struct {
	Version string `json:"version"`
	URL     string `json:"url"`
}

type Settings struct {
	RPCEnabled bool `json:"rpc_enabled"`
}

const settingsFileName = "settings.json"

func checkForUpdate() (*VersionInfo, error) {
	resp, err := http.Get(updateCheckURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch version info: %s", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var versionInfo VersionInfo
	err = json.Unmarshal(data, &versionInfo)
	if err != nil {
		return nil, err
	}

	return &versionInfo, nil
}

func promptForUpdate(versionInfo *VersionInfo, myWindow fyne.Window) {
	parsedURL, err := url.Parse(versionInfo.URL) // Parse the URL string
	if err != nil {
		dialog.ShowError(fmt.Errorf("invalid update URL: %s", versionInfo.URL), myWindow)
		return
	}

	dialog.ShowConfirm(
		"Update Available",
		fmt.Sprintf("A new version (%s) is available. Update now?", versionInfo.Version),
		func(confirmed bool) {
			if confirmed {
				err := fyne.CurrentApp().OpenURL(parsedURL) // Pass the *url.URL type here
				if err != nil {
					dialog.ShowError(err, myWindow)
				}
			}
		},
		myWindow,
	)
}

func loadSettings() (*Settings, error) {
	settingsPath := filepath.Join(os.Getenv("LOCALAPPDATA"), "GoDiag", settingsFileName)
	file, err := os.Open(settingsPath)
	if err != nil {
		if os.IsNotExist(err) {
			return &Settings{RPCEnabled: true}, nil
		}
		return nil, err
	}
	defer file.Close()

	settings := &Settings{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(settings)
	if err != nil {
		return nil, err
	}

	return settings, nil
}

func saveSettings(settings *Settings) error {
	settingsDir := filepath.Join(os.Getenv("LOCALAPPDATA"), "GoDiag")
	if err := os.MkdirAll(settingsDir, os.ModePerm); err != nil {
		return err
	}

	settingsPath := filepath.Join(settingsDir, settingsFileName)
	file, err := os.Create(settingsPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(settings)
}

func main() {
	myApp := app.NewWithID("tv.lewdlilly.GoDiag")
	myWindow := myApp.NewWindow("GoDiag by LewdLillyVT")
	myWindow.Resize(fyne.NewSize(400, 600))

	// Ensure the output directory exists
	outputDir, err := modules.EnsureOutputDir()
	if err != nil {
		dialog.ShowError(err, myWindow)
		return
	}

	// Check for updates
	go func() {
		versionInfo, err := checkForUpdate()
		if err == nil && versionInfo.Version > currentVersion {
			promptForUpdate(versionInfo, myWindow)
		}
	}()

	settings, err := loadSettings()
	if err != nil {
		dialog.ShowError(err, myWindow)
		return
	}

	// Track RPC state
	var rpcRunning = settings.RPCEnabled

	// Start RPC if enabled
	if rpcRunning {
		go func() {
			rpcErr := rpc.StartRPC()
			if rpcErr != nil {
				fyne.CurrentApp().SendNotification(&fyne.Notification{
					Title:   "RPC Error",
					Content: rpcErr.Error(),
				})
			}
		}()
	}

	cleanup := func() {
		if rpcRunning {
			rpc.StopRPC()
		}
	}

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChannel
		cleanup()
		os.Exit(0)
	}()

	myWindow.SetOnClosed(func() {
		cleanup()
	})

	// Diagnostics buttons...
	msInfoButton := widget.NewButton("Generate msinfo32.nfo", func() {
		err := modules.GenerateMsinfo32(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "msinfo32.nfo created successfully", myWindow)
		}
	})

	dxdiagButton := widget.NewButton("Generate dxdiag.txt", func() {
		err := modules.GenerateDxdiag(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "dxdiag.txt created successfully", myWindow)
		}
	})

	etlLogButton := widget.NewButton("Generate Event Trace Log (ETL)", func() {
		err := modules.GenerateETLLog(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "ETL log generated successfully", myWindow)
		}
	})

	biosReportButton := widget.NewButton("Generate BIOS/UEFI Version Report", func() {
		err := modules.GenerateBIOSReport(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "BIOS_Report.txt created successfully", myWindow)
		}
	})

	sysInfoButton := widget.NewButton("Generate Quick System Info", func() {
		err := modules.GenerateQuickSysInfo(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "Quick_System_Info.txt created successfully", myWindow)
		}
	})

	eventLogButton := widget.NewButton("Dump Latest Event Logs", func() {
		err := modules.DumpEventLogs(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "Event_Log_Dump.txt created successfully", myWindow)
		}
	})

	healthButton := widget.NewButton("Generate Drive Health Report", func() {
		err := modules.GenerateHealthAndUsageReport(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "Health_Report.txt created successfully", myWindow)
		}
	})

	securityLogsButton := widget.NewButtonWithIcon("Generate Security and Antivirus Logs", theme.GridIcon(), func() {
		err := modules.GenerateSecurityAndAntivirusLogs(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "Security_Antivirus_Logs.txt created successfully", myWindow)
		}
	})

	networkDiagButton := widget.NewButton("Generate Network Diagnostics Report", func() {
		err := modules.GenerateNetworkReport(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "Network_Diagnostics_Report.txt created successfully", myWindow)
		}
	})

	flushDNSButton := widget.NewButton("Flush DNS Cache", func() {
		err := modules.FlushDNSCache()
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "DNS cache flushed successfully.", myWindow)
		}
	})

	softwareDiagButton := widget.NewButton("Generate Software & Application Report", func() {
		err := modules.GenerateSoftwareReport(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "Software_Diagnostics_Report.txt created successfully", myWindow)
		}
	})

	hardwareButton := widget.NewButton("Generate Hardware & Peripherals Report", func() {
		err := modules.GenerateHardwareReport(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "Hardware_Peripherals_Report.txt created successfully", myWindow)
		}
	})

	driverManagementButton := widget.NewButton("Generate Driver Report", func() {
		err := modules.GenerateDriverReport(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "Driver_Report.txt created successfully", myWindow)
		}
	})

	registryExportButton := widget.NewButton("Export Common Registry Keys", func() {
		err := modules.GenerateRegistryExport(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "Common registry keys exported successfully.\nSee Registry_Export_Summary.txt for details on exported files.", myWindow)
		}
	})

	startupProgramsButton := widget.NewButton("Generate Startup Programs Report", func() {
		err := modules.GenerateStartupProgramsReport(outputDir) // Call the new function
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "Startup_Programs_Report.txt created successfully", myWindow)
		}
	})

	runningProcessesButton := widget.NewButton("Generate Running Processes Report", func() {
		err := modules.GenerateRunningProcessesReport(outputDir) // Call the new function
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "Running_Processes_Report.txt created successfully", myWindow)
		}
	})

	// Help Tab
	linkURL := &url.URL{
		Scheme: "https",
		Host:   "github.com",
		Path:   "LewdLillyVT/godiag",
	}

	helpTab := container.NewTabItem("Help",
		container.NewVBox(
			widget.NewLabel("Help & Documentation"),
			widget.NewLabel("For more information, visit:"),
			widget.NewHyperlink("GoDiag Repository", linkURL),
		),
	)

	// Settings Tab
	rpcToggle := widget.NewCheck("Enable RPC", func(checked bool) {
		rpcRunning = checked
		settings.RPCEnabled = checked
		if err := saveSettings(settings); err != nil {
			dialog.ShowError(err, myWindow)
		}

		if checked {
			go func() {
				rpcErr := rpc.StartRPC()
				if rpcErr != nil {
					fyne.CurrentApp().SendNotification(&fyne.Notification{
						Title:   "RPC Error",
						Content: rpcErr.Error(),
					})
				}
			}()
		} else {
			rpc.StopRPC()
		}
	})
	rpcToggle.SetChecked(settings.RPCEnabled)

	settingsTab := container.NewTabItem("Settings",
		container.NewVBox(
			widget.NewLabel("Settings"),
			rpcToggle,
		),
	)

	// Diagnostics/Main Tab
	mainTab := container.NewTabItem("Main",
		container.NewVBox(
			msInfoButton,
			dxdiagButton,
			sysInfoButton,
			eventLogButton,
			healthButton,
			etlLogButton,
			biosReportButton,
			securityLogsButton,
			networkDiagButton,
			flushDNSButton,
			softwareDiagButton,
			hardwareButton,
			driverManagementButton,
			registryExportButton,
			startupProgramsButton,
			runningProcessesButton,
		),
	)

	tabs := container.NewAppTabs(
		mainTab,
		helpTab,
		settingsTab,
	)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
