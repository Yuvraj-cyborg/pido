package compressors

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"

	"github.com/Yuvraj-cyborg/pido/internal/types"
)

// CompressText gzips the text file and writes "<name>-pido.txt.gz".
// and returns (originalSize, compressedSize, error).
func CompressText(path string, opts types.CompressionOptions) (int64, int64, error) {
	in, err := os.Open(path)
	if err != nil {
		return 0, 0, err
	}
	defer in.Close()

	info, err := in.Stat()
	if err != nil {
		return 0, 0, err
	}
	origSize := info.Size()

	// Building .gz output
	baseOut := buildOutputPath(path, opts.OutDir)
	outPath := baseOut + ".gz"

	out, err := os.Create(outPath)
	if err != nil {
		return 0, 0, err
	}
	defer out.Close()

	gz, err := gzip.NewWriterLevel(out, gzip.BestCompression)
	if err != nil {
		return 0, 0, fmt.Errorf("gzip writer: %w", err)
	}
	defer gz.Close()

	if _, err := io.Copy(gz, in); err != nil {
		return 0, 0, fmt.Errorf("gzip copy: %w", err)
	}

	if err := gz.Close(); err != nil {
		return origSize, 0, err
	}
	if err := out.Sync(); err != nil {
		return origSize, 0, err
	}

	outInfo, err := os.Stat(outPath)
	if err != nil {
		return origSize, 0, err
	}
	return origSize, outInfo.Size(), nil
}
