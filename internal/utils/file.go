package utils

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/Yuvraj-cyborg/pido/internal/types"
)

var supportedExt = map[string]struct{}{
	".jpg": {}, ".jpeg": {}, ".png": {},
	".pdf": {},
	".txt": {},
}

func collectFiles(opts types.CompressionOptions) []string {
	var results []string

	if len(opts.Files) > 0 {
		for _, f := range opts.Files {
			if exists, _ := fileExists(f); !exists {
				continue
			}
			if isSupported(f) {
				results = append(results, f)
			}
		}
		return results
	}

	if opts.Dir != "" {
		filepath.WalkDir(opts.Dir, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if d.IsDir() {
				return nil
			}
			if isSupported(path) {
				results = append(results, path)
			}
			return nil
		})
	}

	return results
}

func isSupported(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	_, ok := supportedExt[ext]
	return ok
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}

// BuildOutputPath returns path with "-pido" suffix placed next to input
// or inside outDir if provided.
func BuildOutputPath(inputPath, outDir string) (string, error) {
	base := filepath.Base(inputPath)
	ext := filepath.Ext(base)
	name := strings.TrimSuffix(base, ext)
	newName := name + "-pido" + ext

	outFolder := filepath.Dir(inputPath)
	if outDir != "" {
		outFolder = outDir
		if err := EnsureDir(outFolder); err != nil {
			return "", err
		}
	}
	return filepath.Join(outFolder, newName), nil
}

func EnsureDir(dir string) error {
	return os.MkdirAll(dir, 0o755)
}
