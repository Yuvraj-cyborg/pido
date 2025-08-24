package compressors

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/Yuvraj-cyborg/pido/internal/types"
	"golang.org/x/image/draw"
)

// CompressImage returns (originalSize, compressedSize, error)
// JPEG: re-encode with quality.
// PNG: lossless best-compression; if quality<100, downscale proportionally to reduce size.
func CompressImage(path string, opts types.CompressionOptions) (int64, int64, error) {
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

	img, format, err := image.Decode(in)
	if err != nil {
		return 0, 0, fmt.Errorf("decode: %w", err)
	}

	// For PNG, we downscale to approximate "quality" intent.
	// And for JPEG, we rely primarily on quality re-encode (can also downscale).
	if strings.EqualFold(format, "png") && opts.Quality < 100 {
		scale := float64(opts.Quality) / 100.0
		if scale < 0.1 {
			scale = 0.1
		}
		b := img.Bounds()
		newW := int(float64(b.Dx()) * scale)
		newH := int(float64(b.Dy()) * scale)
		if newW >= 1 && newH >= 1 {
			dst := image.NewRGBA(image.Rect(0, 0, newW, newH))
			draw.CatmullRom.Scale(dst, dst.Bounds(), img, b, draw.Over, nil)
			img = dst
		}
	}

	outPath := buildOutputPath(path, opts.OutDir)

	out, err := os.Create(outPath)
	if err != nil {
		return 0, 0, err
	}
	defer out.Close()

	switch strings.ToLower(filepath.Ext(path)) {
	case ".jpg", ".jpeg":
		q := opts.Quality
		if q < 1 {
			q = 1
		}
		if q > 100 {
			q = 100
		}
		if err := jpeg.Encode(out, img, &jpeg.Options{Quality: q}); err != nil {
			return 0, 0, fmt.Errorf("jpeg encode: %w", err)
		}
	case ".png":
		enc := png.Encoder{CompressionLevel: png.BestCompression}
		if err := enc.Encode(out, img); err != nil {
			return 0, 0, fmt.Errorf("png encode: %w", err)
		}
	default:
		return 0, 0, fmt.Errorf("unsupported image extension")
	}

	newInfo, err := out.Stat()
	if err != nil {
		return origSize, 0, err
	}
	return origSize, newInfo.Size(), nil
}
