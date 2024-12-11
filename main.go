package main

import (
	"bytes"
	_ "embed" // Used for embedding files
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// Embed the vrpathreg.exe file & openvr_api.dll
//
//go:embed assets/vrpathreg.exe
var embeddedVrPathReg []byte

//go:embed assets/openvr_api.dll
var embeddedOpenVRAPI []byte

func main() {
	// Initialize the Fyne app
	myApp := app.NewWithID("tv.lewdlilly.GoDiag")
	myWindow := myApp.NewWindow("System Diagnostics Tool by LewdLillyVT")
	myWindow.Resize(fyne.NewSize(400, 600))

	// Define the output directory for diagnostic files
	outputDir := filepath.Join(os.TempDir(), "DiagnosticsFiles")
	os.MkdirAll(outputDir, os.ModePerm)

	// Buttons to generate diagnostics
	msInfoButton := widget.NewButton("Generate msinfo32.nfo", func() {
		err := generateMsinfo32(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "msinfo32.nfo created successfully", myWindow)
		}
	})

	dxdiagButton := widget.NewButton("Generate dxdiag.txt", func() {
		err := generateDxdiag(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "dxdiag.txt created successfully", myWindow)
		}
	})

	sysInfoButton := widget.NewButton("Generate Quick System Info", func() {
		err := generateQuickSysInfo(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "Quick_System_Info.txt created successfully", myWindow)
		}
	})

	eventLogButton := widget.NewButton("Dump Latest Event Logs", func() {
		err := dumpEventLogs(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "Event_Log_Dump.txt created successfully", myWindow)
		}
	})

	healthButton := widget.NewButton("Generate Drive Health Report", func() {
		err := generateHealthAndUsageReport(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "Health_Report.txt created successfully", myWindow)
		}
	})

	steamVRButton := widget.NewButton("Generate SteamVR Report", func() {
		err := generateSteamVRReport(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "SteamVR_Report.txt created successfully", myWindow)
		}
	})

	// Adding ETL logs button
	etlLogButton := widget.NewButton("Generate Event Trace Log (ETL)", func() {
		err := generateETLLog(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "ETL log generated successfully", myWindow)
		}
	})

	// Adding Security and Antivirus Logs button
	securityLogsButton := widget.NewButtonWithIcon("Generate Security and Antivirus Logs", theme.GridIcon(), func() {
		err := generateSecurityAndAntivirusLogs(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "Security_Antivirus_Logs.txt created successfully", myWindow)
		}
	})

	// Adding BIOS/UEFI Version Report button
	biosReportButton := widget.NewButton("Generate BIOS/UEFI Version Report", func() {
		err := generateBIOSReport(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "BIOS_Report.txt created successfully", myWindow)
		}
	})

	// Layout all buttons in a vertical box
	content := container.NewVBox(
		widget.NewLabel("Diagnostics Tool"),
		widget.NewLabel("Files saved to: "+outputDir),
		msInfoButton,
		dxdiagButton,
		sysInfoButton,
		eventLogButton,
		healthButton,
		steamVRButton,
		etlLogButton,
		securityLogsButton,
		biosReportButton,
	)

	// Set up the main window content
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

// generateMsinfo32 generates msinfo32.nfo file and saves it to the output directory.
func generateMsinfo32(outputDir string) error {
	outputPath := filepath.Join(outputDir, "msinfo32.nfo")
	cmd := exec.Command("msinfo32", "/nfo", outputPath)
	return cmd.Run()
}

// generateDxdiag generates dxdiag.txt file and saves it to the output directory.
func generateDxdiag(outputDir string) error {
	outputPath := filepath.Join(outputDir, "dxdiag.txt")
	cmd := exec.Command("dxdiag", "/t", outputPath)
	return cmd.Run()
}

// generateQuickSysInfo gathers basic system information like CPU, GPU, RAM, and saves it to a text file.
func generateQuickSysInfo(outputDir string) error {
	var output bytes.Buffer
	outputPath := filepath.Join(outputDir, "Quick_System_Info.txt")

	// Collect CPU information
	output.WriteString("CPU Information:\n")
	cpuInfo, err := exec.Command("wmic", "cpu", "get", "Name,MaxClockSpeed,Manufacturer").Output()
	if err != nil {
		output.WriteString("Error gathering CPU information.\n")
	} else {
		output.Write(cpuInfo)
	}

	// Collect GPU information
	output.WriteString("\nGPU Information:\n")
	gpuInfo, err := exec.Command("wmic", "path", "win32_videocontroller", "get", "name,driverversion").Output()
	if err != nil {
		output.WriteString("Error gathering GPU information.\n")
	} else {
		output.Write(gpuInfo)
	}

	// Collect RAM information
	output.WriteString("\nRAM Information:\n")
	ramInfo, err := exec.Command("wmic", "memorychip", "get", "capacity,manufacturer,partnumber,speed").Output()
	if err != nil {
		output.WriteString("Error gathering RAM information.\n")
	} else {
		output.Write(ramInfo)
	}

	// Collect Motherboard and BIOS information
	output.WriteString("\nMotherboard and BIOS Information:\n")
	moboInfo, err := exec.Command("wmic", "baseboard", "get", "product,manufacturer").Output()
	if err != nil {
		output.WriteString("Error gathering Motherboard information.\n")
	} else {
		output.Write(moboInfo)
	}
	biosInfo, err := exec.Command("wmic", "bios", "get", "version,serialnumber").Output()
	if err != nil {
		output.WriteString("Error gathering BIOS information.\n")
	} else {
		output.Write(biosInfo)
	}

	// Additional system details if on Windows
	if runtime.GOOS == "windows" {
		output.WriteString("\nOS Information:\n")
		osInfo, err := exec.Command("systeminfo").Output()
		if err != nil {
			output.WriteString("Error gathering OS information.\n")
		} else {
			output.Write(osInfo)
		}
	}

	// Append footer
	output.WriteString("\n\nReport generated by GoDiag. Learn more at https://github.com/LewdLillyVT/godiag")

	// Save to file
	return os.WriteFile(outputPath, output.Bytes(), 0644)
}

// dumpEventLogs extracts the last 10 warnings, errors, and critical errors from the event logs.
func dumpEventLogs(outputDir string) error {
	var output bytes.Buffer
	outputPath := filepath.Join(outputDir, "Event_Log_Dump.txt")

	// Gather last 10 warnings
	output.WriteString("Last 10 Warning Events:\n")
	warnings, err := exec.Command("wevtutil", "qe", "System", "/q:*[System[(Level=3)]]", "/c:10", "/f:text").Output()
	if err != nil {
		output.WriteString("Error gathering warning events.\n")
	} else {
		output.Write(warnings)
	}

	// Gather last 10 errors
	output.WriteString("\nLast 10 Error Events:\n")
	errors, err := exec.Command("wevtutil", "qe", "System", "/q:*[System[(Level=2)]]", "/c:10", "/f:text").Output()
	if err != nil {
		output.WriteString("Error gathering error events.\n")
	} else {
		output.Write(errors)
	}

	// Gather last 10 critical errors
	output.WriteString("\nLast 10 Critical Events:\n")
	criticals, err := exec.Command("wevtutil", "qe", "System", "/q:*[System[(Level=1)]]", "/c:10", "/f:text").Output()
	if err != nil {
		output.WriteString("Error gathering critical events.\n")
	} else {
		output.Write(criticals)
	}

	// Append footer
	output.WriteString("\n\nReport generated by GoDiag. Learn more at https://github.com/LewdLillyVT/godiag")

	// Save to file
	return os.WriteFile(outputPath, output.Bytes(), 0644)
}

// generateHealthAndUsageReport gathers detailed health information of drives.
func generateHealthAndUsageReport(outputDir string) error {
	var output bytes.Buffer
	outputPath := filepath.Join(outputDir, "Health_Report.txt")

	// Collect Drive Health Information
	output.WriteString("Drive Health Information:\n")

	// Retrieve detailed drive information including model, serial number, size, and status
	driveInfo, err := exec.Command("wmic", "diskdrive", "get", "Model,SerialNumber,Size,Status").Output()
	if err != nil {
		output.WriteString("Error gathering drive health information.\n")
	} else {
		output.Write(driveInfo)
	}

	// Collect SMART data if available for each drive
	output.WriteString("\nSMART Data for Drives:\n")
	smartData, err := exec.Command("wmic", "diskdrive", "get", "Status,LastErrorCode,Capabilities,CapabilityDescriptions").Output()
	if err != nil {
		output.WriteString("Error gathering SMART data.\n")
	} else {
		output.Write(smartData)
	}

	// Append footer
	output.WriteString("\n\nReport generated by GoDiag. Learn more at https://github.com/LewdLillyVT/godiag")

	// Save to file
	return os.WriteFile(outputPath, output.Bytes(), 0644)
}

// extractVrPathReg extracts the embedded vrpathreg.exe and openvr_api.dll to a temporary directory
func extractVrPathReg() (string, error) {
	tempDir := os.TempDir()
	vrpathregPath := filepath.Join(tempDir, "vrpathreg.exe")
	openvrAPIPath := filepath.Join(tempDir, "openvr_api.dll")

	// Write vrpathreg.exe to the temporary directory
	err := os.WriteFile(vrpathregPath, embeddedVrPathReg, 0755)
	if err != nil {
		return "", err
	}

	// Write openvr_api.dll to the temporary directory
	err = os.WriteFile(openvrAPIPath, embeddedOpenVRAPI, 0644)
	if err != nil {
		return "", err
	}

	return vrpathregPath, nil
}

// generateSteamVRReport generates a SteamVR diagnostic report and saves it to the output directory.
func generateSteamVRReport(outputDir string) error {
	// Define the path for the SteamVR report
	outputPath := filepath.Join(outputDir, "SteamVR_Report.txt")

	// Extract the embedded files (vrpathreg.exe and openvr_api.dll)
	vrpathregPath, err := extractVrPathReg()
	if err != nil {
		return err
	}
	defer os.Remove(vrpathregPath) // Clean up vrpathreg.exe after use

	openvrAPIPath := filepath.Join(os.TempDir(), "openvr_api.dll")
	defer os.Remove(openvrAPIPath) // Clean up openvr_api.dll after use

	// Attempt to gather SteamVR logs from the Steam logs directory
	steamLogDirectory := filepath.Join(os.Getenv("ProgramFiles(x86)"), "Steam", "logs")
	if steamLogDirectory == "" {
		steamLogDirectory = filepath.Join(os.Getenv("ProgramFiles"), "Steam", "logs")
	}

	// Create a buffer to hold the output report
	var output bytes.Buffer
	output.WriteString("SteamVR Logs:\n")

	// Check if Steam logs exist in the default directory
	files, err := os.ReadDir(steamLogDirectory)
	if err != nil {
		output.WriteString("Error reading Steam logs directory: " + err.Error() + "\n")
	} else if len(files) == 0 {
		output.WriteString("No log files found in Steam's default log directory.\n")
	} else {
		// If logs are found, write the paths of the logs to the report
		for _, file := range files {
			if filepath.Ext(file.Name()) == ".log" && filepath.Base(file.Name())[:9] == "vrmonitor" {
				logFilePath := filepath.Join(steamLogDirectory, file.Name())
				output.WriteString("Log file path: " + logFilePath + "\n")
			}
		}
	}

	// If logs are not found, attempt to generate logs using vrpathreg
	if len(output.Bytes()) == len("SteamVR Logs:\n") {
		output.WriteString("\nLogs not found in Steam's default log directory. Generating new logs.\n")

		// Generate the SteamVR path information (but don't save it to a separate log file)
		cmd := exec.Command(vrpathregPath, "show")
		var cmdOutput bytes.Buffer
		cmd.Stdout = &cmdOutput
		cmd.Stderr = &cmdOutput

		err := cmd.Run()
		if err != nil {
			output.WriteString("Error generating SteamVR path info:\n")
			output.WriteString(cmdOutput.String())
		} else {
			output.WriteString("\nSteamVR Path Information:\n")
			output.Write(cmdOutput.Bytes())
		}
	}

	// Append the message about the functionality being worked on
	output.WriteString("\n\nNote: The functionality to get the logs and save them in the %LOCALAPPDATA%\\Temp\\DiagnosticsFiles directory is currently being worked on.\n")

	// Save the SteamVR report to the specified output path (in the outputDir)
	err = os.WriteFile(outputPath, output.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

// generateETLLog generates an Event Trace Log (ETL) file and saves it to the output directory.
func generateETLLog(outputDir string) error {
	outputPath := filepath.Join(outputDir, "event_trace_log.etl")

	// Use wevtutil to export the system log to an ETL file
	cmd := exec.Command("wevtutil", "epl", "System", outputPath)
	var cmdOutput bytes.Buffer
	cmd.Stdout = &cmdOutput
	cmd.Stderr = &cmdOutput

	err := cmd.Run()
	if err != nil {
		return err
	}

	// If successful, we can optionally log the success message to the output
	return nil
}

// generateSecurityAndAntivirusLogs extracts recent security and antivirus-related events.
func generateSecurityAndAntivirusLogs(outputDir string) error {
	// Ensure the application has administrative privileges
	if err := ensureAdminPrivileges(); err != nil {
		return err
	}

	var output bytes.Buffer
	outputPath := filepath.Join(outputDir, "Security_Antivirus_Logs.txt")

	// Gather Security logs
	output.WriteString("Security Logs:\n")
	securityLogs, err := exec.Command("wevtutil", "qe", "Security", "/c:50", "/f:text").Output()
	if err != nil {
		output.WriteString("Error gathering security logs.\n")
	} else {
		output.Write(securityLogs)
	}

	// Gather Antivirus logs (specific to Windows Defender)
	output.WriteString("\nWindows Defender Logs:\n")
	antivirusLogs, err := exec.Command("powershell", "Get-MpThreatDetection").Output()
	if err != nil {
		output.WriteString("Error gathering antivirus logs.\n")
	} else {
		output.Write(antivirusLogs)
	}

	// Append footer
	output.WriteString("\n\nReport generated by GoDiag. Learn more at https://github.com/LewdLillyVT/godiag")

	// Save to file
	return os.WriteFile(outputPath, output.Bytes(), 0644)
}

// ensureAdminPrivileges restarts the application with administrative privileges if not already elevated.
func ensureAdminPrivileges() error {
	if runtime.GOOS != "windows" {
		return nil // Elevation is only required for Windows
	}

	// Check if the current process is already running with administrative privileges
	isAdmin, err := isRunningAsAdmin()
	if err != nil {
		return fmt.Errorf("failed to check admin privileges: %v", err)
	}
	if isAdmin {
		return nil // Already elevated, no need to restart
	}

	// Determine the path of the executable
	exePath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to find executable path: %v", err)
	}

	// If running from a source file (like during testing), ensure the path resolves correctly
	exePath, err = filepath.EvalSymlinks(exePath)
	if err != nil {
		return fmt.Errorf("failed to resolve executable symlinks: %v", err)
	}

	// Use PowerShell to relaunch the application with administrative privileges
	cmd := exec.Command("powershell", "-Command", "Start-Process", "\""+exePath+"\"", "-Verb", "runAs")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start process with elevated privileges: %v", err)
	}

	// Exit the current process since the elevated instance takes over
	os.Exit(0)
	return nil
}

// isRunningAsAdmin checks if the current process is running with administrative privileges.
func isRunningAsAdmin() (bool, error) {
	cmd := exec.Command("powershell", "-Command", "([Security.Principal.WindowsPrincipal] [Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)")
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	// Check if the output is "True"
	return bytes.Equal(bytes.TrimSpace(output), []byte("True")), nil
}

// generateBIOSReport gathers BIOS/UEFI version information and saves it to a text file.
func generateBIOSReport(outputDir string) error {
	var output bytes.Buffer
	outputPath := filepath.Join(outputDir, "BIOS_Report.txt")

	output.WriteString("BIOS/UEFI Version Information:\n")
	biosInfo, err := exec.Command("wmic", "bios", "get", "Manufacturer,SMBIOSBIOSVersion,ReleaseDate").Output()
	if err != nil {
		output.WriteString("Error gathering BIOS information.\n")
	} else {
		output.Write(biosInfo)
	}

	// Append footer
	output.WriteString("\n\nReport generated by GoDiag. Learn more at https://github.com/LewdLillyVT/godiag")

	// Save to file
	return os.WriteFile(outputPath, output.Bytes(), 0644)
}
