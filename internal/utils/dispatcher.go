package utils

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Yuvraj-cyborg/pido/internal/compressors"
	"github.com/Yuvraj-cyborg/pido/internal/types"
)

func DispatchCompression(opts types.CompressionOptions) error {
	if opts.Quality < 0 || opts.Quality > 100 {
		return fmt.Errorf("quality must be between 0 and 100")
	}

	files := collectFiles(opts)
	if len(files) == 0 {
		return fmt.Errorf("no supported files found (jpg/jpeg/png/pdf/txt)")
	}

	var totalBefore, totalAfter int64
	for _, f := range files {
		ext := strings.ToLower(filepath.Ext(f))
		var before, after int64
		var err error

		switch ext {
		case ".jpg", ".jpeg", ".png":
			before, after, err = compressors.CompressImage(f, opts)
		case ".pdf":
			before, after, err = compressors.CompressPDF(f, opts.Quality, opts.OutDir)
		case ".txt":
			before, after, err = compressors.CompressText(f, opts)
		default:
			continue
		}
		if err != nil {
			return fmt.Errorf("compress %s: %w", f, err)
		}

		totalBefore += before
		totalAfter += after

		savePct := pctSaved(before, after)
		fmt.Printf("%s → %d KB → %d KB (%d%% saved)\n", f, before/1024, after/1024, savePct)
	}

	fmt.Printf("\nTotal: %d KB → %d KB (%d%% saved)\n",
		totalBefore/1024, totalAfter/1024, pctSaved(totalBefore, totalAfter))
	return nil
}

func pctSaved(before, after int64) int64 {
	if before <= 0 {
		return 0
	}
	if after < 0 {
		after = 0
	}
	return (before - after) * 100 / before
}
