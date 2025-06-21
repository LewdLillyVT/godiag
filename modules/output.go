package modules

import (
	"os"
	"path/filepath"
)

// customOutputDir stores the user-selected output directory. If empty, the default path is used.
var customOutputDir string

// SetCustomOutputDir sets the custom output directory. This function should be called
// when the user selects a new directory via the settings.
func SetCustomOutputDir(path string) {
	customOutputDir = path
}

// GetCustomOutputDir returns the currently set custom output directory.
// This can be useful for displaying the selected path in the UI.
func GetCustomOutputDir() string {
	return customOutputDir
}

// EnsureOutputDir creates the output directory and returns its path.
// It uses the customOutputDir if it has been set; otherwise, it defaults
// to `%LOCALAPPDATA%\Temp\DiagnosticsFiles`.
func EnsureOutputDir() (string, error) {
	var targetDir string
	if customOutputDir != "" {
		targetDir = customOutputDir
	} else {
		targetDir = filepath.Join(os.Getenv("LOCALAPPDATA"), "Temp", "DiagnosticsFiles")
	}

	// Create all necessary parent directories if they don't exist
	err := os.MkdirAll(targetDir, os.ModePerm)
	if err != nil {
		return "", err
	}
	return targetDir, nil
}
