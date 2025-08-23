package utils

import (
	"fmt"

	"github.com/Yuvraj-cyborg/pido/internal/types"
)

func PrintReport(results []types.FileInfo) {
	fmt.Println("Compression Summary:")
	for _, r := range results {
		saving := 100 - (100 * r.CompressedSize / r.OriginalSize)
		fmt.Printf("%s → %s | %d KB → %d KB (%d%% saved)\n",
			r.Path,
			r.Type,
			r.OriginalSize/1024,
			r.CompressedSize/1024,
			saving,
		)
	}
}
