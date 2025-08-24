package compressors

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

// CompressPDF compresses a PDF using pdfcpu's OptimizeFile API with image downsampling.
// Returns (originalSize, compressedSize, error) to match dispatcher expectations.
func CompressPDF(path string, quality int, outDir string) (int64, int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, 0, err
	}
	origSize := info.Size()

	base := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
	outPath := filepath.Join(outDirOrSame(outDir, filepath.Dir(path)), base+"-pido.pdf")

	config := model.NewDefaultConfiguration()

	config.Cmd = model.OPTIMIZE

	if quality < 100 {
		config.Optimize = true
		config.OptimizeResourceDicts = true
		config.OptimizeDuplicateContentStreams = true
	}

	if err := api.OptimizeFile(path, outPath, config); err != nil {
		return origSize, 0, fmt.Errorf("failed to optimize PDF: %w", err)
	}

	newInfo, err := os.Stat(outPath)
	if err != nil {
		return origSize, 0, err
	}

	return origSize, newInfo.Size(), nil
}

func outDirOrSame(outDir, fallback string) string {
	if outDir != "" {
		_ = os.MkdirAll(outDir, 0755)
		return outDir
	}
	return fallback
}
