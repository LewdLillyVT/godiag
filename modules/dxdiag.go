package modules

import (
	"os/exec"
	"path/filepath"
)

func GenerateDxdiag(outputDir string) error {
	outputPath := filepath.Join(outputDir, "dxdiag.txt")
	cmd := exec.Command("dxdiag", "/t", outputPath)
	return cmd.Run()
}
