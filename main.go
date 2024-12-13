package main

import (
	"GoDiag-beta/modules"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.NewWithID("tv.lewdlilly.GoDiag.beta")
	myWindow := myApp.NewWindow("GoDiag Beta by LewdLillyVT")
	myWindow.Resize(fyne.NewSize(400, 600))

	// Ensure the output directory exists
	outputDir, err := modules.EnsureOutputDir()
	if err != nil {
		dialog.ShowError(err, myWindow)
		return
	}

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

	// Adding ETL logs button
	etlLogButton := widget.NewButton("Generate Event Trace Log (ETL)", func() {
		err := modules.GenerateETLLog(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "ETL log generated successfully", myWindow)
		}
	})

	// Adding BIOS/UEFI Version Report button
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

	// Adding Security and Antivirus Logs button
	securityLogsButton := widget.NewButtonWithIcon("Generate Security and Antivirus Logs", theme.GridIcon(), func() {
		err := modules.GenerateSecurityAndAntivirusLogs(outputDir)
		if err != nil {
			dialog.ShowError(err, myWindow)
		} else {
			dialog.ShowInformation("Success", "Security_Antivirus_Logs.txt created successfully", myWindow)
		}
	})

	// Add other buttons here following the same pattern...

	content := container.NewVBox(
		widget.NewLabel("Diagnostics Tool"),
		widget.NewLabel("Files saved to: "+outputDir),
		msInfoButton,
		dxdiagButton,
		sysInfoButton,
		eventLogButton,
		healthButton,
		etlLogButton,
		biosReportButton,
		securityLogsButton,
		// Add other buttons here
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
