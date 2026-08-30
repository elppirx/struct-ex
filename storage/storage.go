package storage

import (
	"encoding/json"
	"os"
)

func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func IsJSON(data []byte) bool {
	return json.Valid(data)
}
