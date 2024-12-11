# GoDiag

GoDiag is a powerful diagnostic tool that collects essential system information, providing comprehensive reports and summaries. This utility generates `.nfo`, `.etl` and `.txt` reports, dumps the latest system events, and offers a detailed system info report for troubleshooting and system analysis.

## Features

- **MSInfo Report**: Generate a detailed `msinfo32.nfo` file, containing an extensive overview of system configuration and components.
- **DxDiag Report**: Create a `dxdiag.txt` file, which provides diagnostic information on DirectX components and drivers, valuable for troubleshooting graphical or hardware issues.
- **System Events Dump**: Capture and output recent system event logs, including errors, warnings, and critical events.
- **Comprehensive System Report**: Generate a summary report with core system information in one place.
- **SteamVR Report**: Generate a SteamVR diagnostic report. *(Work in Progress)*
- **ETL Logs**: Extract Event Trace Log (ETL) files from the system for detailed event analysis.

## Preview

Below is a preview of the GoDiag interface and output:

![GoDiag Preview](https://cdn.hyrule.pics/52b31a0cd.png)

## Usage

To run GoDiag, simply run the exe provided in the releases or build it yourself.

## Output

GoDiag will generate the following files in the `/LOCALAPPDATA/Temp/DiagnosticsFiles` directory:

- **msinfo32.nfo** - A detailed system configuration file.
- **dxdiag.txt** - A DirectX diagnostics report.
- **events.log** - Recent system events.
- **system_info_report.txt** - A summary report of core system information.
- **Health_Report.txt** - A summary of your drive health.
- **SteamVR_Report.txt** - SteamVR diagnostic report. *(Work in Progress)*
- **event_trace_log.etl** - Extracted Event Trace Logs for system events and performance monitoring.

