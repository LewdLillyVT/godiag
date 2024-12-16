package modules

import (
	"bytes"
	"os/exec"
	"path/filepath"
)

// generateETLLog generates an Event Trace Log (ETL) file and saves it to the output directory.
func GenerateETLLog(outputDir string) error {
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
