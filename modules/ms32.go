package modules

import (
	"os/exec"
	"path/filepath"
)

func GenerateMsinfo32(outputDir string) error {
	outputPath := filepath.Join(outputDir, "msinfo32.nfo")
	cmd := exec.Command("msinfo32", "/nfo", outputPath)
	return cmd.Run()
}
