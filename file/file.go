package file

import (
	"encoding/json"
	"os"
	"struct-ex/bins"
)

func CreateFile(b *bins.Bin, filename string) error {
	data, err := json.Marshal(b)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(filename string) (*bins.Bin, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	b := &bins.Bin{}
	err = json.Unmarshal(data, b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
