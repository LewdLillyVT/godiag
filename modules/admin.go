package modules

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

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
