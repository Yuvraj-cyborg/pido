package compressors

import (
	"os"
	"path/filepath"
	"strings"
)

// buildOutputPath creates output path with -pido suffix
func buildOutputPath(inputPath, outDir string) string {
	base := filepath.Base(inputPath)
	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)
	newName := name + "-pido" + ext

	outFolder := filepath.Dir(inputPath)
	if outDir != "" {
		outFolder = outDir
		os.MkdirAll(outFolder, 0755)
	}
	return filepath.Join(outFolder, newName)
}
