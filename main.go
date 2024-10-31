package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Initialize the Fyne app
	myApp := app.New()
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

	// Layout all buttons in a vertical box
	content := container.NewVBox(
		widget.NewLabel("Diagnostics Tool"),
		widget.NewLabel("Files saved to: "+outputDir),
		msInfoButton,
		dxdiagButton,
		sysInfoButton,
		eventLogButton,
		healthButton,
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

	// Save to file
	return os.WriteFile(outputPath, output.Bytes(), 0644)
}
