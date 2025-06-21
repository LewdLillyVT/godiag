# GoDiag

GoDiag is a powerful diagnostic tool that collects essential system information, providing comprehensive reports and summaries. This utility generates `.nfo`, `.evtx`, and `.txt` reports, dumps the latest system events, and offers a detailed system info report for troubleshooting and system analysis.

## Features

-   **MSInfo Report**: Generates a detailed `msinfo32.nfo` file, containing an extensive overview of system configuration and components.
-   **DxDiag Report**: Creates a `dxdiag.txt` file, which provides diagnostic information on DirectX components and drivers, valuable for troubleshooting graphical or hardware issues.
-   **System Events Dump**: Captures and outputs recent system event logs, including errors, warnings, and critical events.
-   **Comprehensive System Report**: Generates a summary report with core system information in one place.
-   **SteamVR Report**: Generates a SteamVR diagnostic report. *(Work in Progress)*
-   **ETL Logs**: Extracts Event Trace Log (ETL) files from the system for detailed event analysis.
-   **System Security & Antivirus Logs**: Extracts the latest system security and Windows Defender logs.
-   **BIOS/UEFI Version Report**: Reads the latest version information of the BIOS/UEFI.
-   **Network Diagnostics**: Exports network diagnostics and allows DNS flushing.
-   **Hardware Info**: Provides information about connected USB devices, printers, and battery health.
-   **Driver Report**: Exports information about installed drivers, including versions and available install dates.
-   **Registry Export**: Exports a list of commonly diagnosed registry keys into a dedicated subfolder.
-   **Startup Programs Report**: Collects and reports on programs configured to run automatically at system startup from various locations.
-   **Running Processes Report**: Provides a detailed list of all processes currently active on the system.

## Preview

Below is a preview of the GoDiag interface and output:

![GoDiag Preview](https://cdn.hyrule.pics/52b31a0cd.png)

## Usage

To run GoDiag, simply execute the `.exe` provided in the releases or build it yourself.

## Output

GoDiag will generate the following files and folders in the `%LOCALAPPDATA%\Temp\DiagnosticsFiles` directory:

-   **msinfo32.nfo**: A detailed system configuration file.
-   **dxdiag.txt**: A DirectX diagnostics report.
-   **Event_Log_Dump.txt**: Recent system events (warnings, errors, critical).
-   **Quick_System_Info.txt**: A summary report of core system information (CPU, GPU, RAM, OS).
-   **Health_Report.txt**: A summary of your drive health and SMART data.
-   **SteamVR_Report.txt**: SteamVR diagnostic report. *(Work in Progress)*
-   **event_trace_log.evtx**: Extracted Event Trace Logs for system events and performance monitoring.
-   **Security_Antivirus_Logs.txt**: A summary of the latest Windows security and Windows Defender logs.
-   **BIOS_Report.txt**: Shows the latest BIOS/UEFI version information.
-   **Network_Diagnostics_Report.txt**: Comprehensive network configuration, connections, and connectivity test.
-   **Hardware_Peripherals_Report.txt**: Information about connected USB devices, printers, and battery health.
-   **Driver_Report.txt**: Lists all installed drivers with verbose details.
-   **Registry_Export_Summary.txt**: A summary of exported registry keys.
-   **RegistryExports/**: A folder containing `.reg` files for commonly diagnosed registry keys.
-   **Startup_Programs_Report.txt**: A report detailing programs configured to run on system startup.
-   **Running_Processes_Report.txt**: A comprehensive list of all currently active processes.
