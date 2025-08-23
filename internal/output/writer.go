package output

import (
	"os"
)

func SaveFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0644)
}
