package file

import (
	"os"
	"path/filepath"
)

type File interface {
	ReadFile(string) (string, error)
	IsJSON(string) bool
}

func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func IsJSON(filename string) bool {
	return filepath.Ext(filename) == ".json"
}
