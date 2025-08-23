package types

type CompressionOptions struct {
	Files   []string
	Dir     string
	Quality int
	OutDir  string
}

type FileInfo struct {
	Path           string
	OriginalSize   int64
	CompressedSize int64
	Type           string
}
