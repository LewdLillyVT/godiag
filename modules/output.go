package modules

import (
	"os"
	"path/filepath"
)

// EnsureOutputDir creates the output directory and returns its path.
func EnsureOutputDir() (string, error) {
	outputDir := filepath.Join(os.Getenv("LOCALAPPDATA"), "Temp", "DiagnosticsFiles", "beta")
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return "", err
	}
	return outputDir, nil
}
