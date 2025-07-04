package modules

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// GenerateRunningProcessesReport collects information about all currently running processes
// and saves it to a text file.
func GenerateRunningProcessesReport(outputDir string) error {
	var output bytes.Buffer
	outputPath := filepath.Join(outputDir, "Running_Processes_Report.txt")

	output.WriteString("--- Running Processes Report ---\n\n")
	output.WriteString("This report lists all processes currently running on the system.\n\n")

	// Use the 'tasklist' command to get a list of running processes.
	// '/v' for verbose output (e.g., session name, PID, memory usage, window title),
	// '/fo list' for a detailed, list-formatted output.
	cmd := exec.Command("tasklist", "/v", "/fo", "list")
	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error running tasklist command: %v\nOutput: %s", err, string(cmdOutput))
	}

	output.Write(cmdOutput)
	output.WriteString("\n")

	// --- Footer ---
	output.WriteString("\n\nReport generated by GoDiag. Learn more at https://github.com/LewdLillyVT/godiag")

	// Save the collected information to the designated output file
	return os.WriteFile(outputPath, output.Bytes(), 0644)
}
