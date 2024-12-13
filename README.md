# Welcome to the GoDiag beta channel
Here in the beta channel you will find all the upcoming features for GoDiag.

To use the beta release you can download the latest pre-release or build the beta release yourself
---

GoDiag is a powerful diagnostic tool that collects essential system information, providing comprehensive reports and summaries. This utility generates `.nfo`, `.etl` and `.txt` reports, dumps the latest system events, and offers a detailed system info report for troubleshooting and system analysis.

## Features

- **MSInfo Report**: Generate a detailed `msinfo32.nfo` file, containing an extensive overview of system configuration and components.
- **DxDiag Report**: Create a `dxdiag.txt` file, which provides diagnostic information on DirectX components and drivers, valuable for troubleshooting graphical or hardware issues.
- **System Events Dump**: Capture and output recent system event logs, including errors, warnings, and critical events.
- **Comprehensive System Report**: Generate a summary report with core system information in one place.
- **ETL Logs**: Extract Event Trace Log (ETL) files from the system for detailed event analysis.
- **System Security & Antivirus Logs**: Extract the latest system security and windows defender logs.
- **BIOS/UEFI Version Report**: Reads the latest version info of the BIOS/UEFI.

## Preview

Below is a preview of the GoDiag interface and output:

![GoDiag Preview](https://cdn.hyrule.pics/52b31a0cd.png)

## Usage

To run GoDiag, simply run the exe provided in the releases or build it yourself.

## Output

GoDiag will generate the following files in the `/LOCALAPPDATA/Temp/DiagnosticsFiles/beta` directory:

- **msinfo32.nfo** - A detailed system configuration file.
- **dxdiag.txt** - A DirectX diagnostics report.
- **events.log** - Recent system events.
- **system_info_report.txt** - A summary report of core system information.
- **Health_Report.txt** - A summary of your drive health.
- **event_trace_log.etl** - Extracted Event Trace Logs for system events and performance monitoring.
- **Security_Antivirus_Logs.txt** - A summary of the latest windows security and windows defender logs.
- **BIOS_Report.txt** - Shows the latest BIOS/UEFI version info.

