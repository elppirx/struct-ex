package storage

import (
	"encoding/json"
	"os"
	"struct-ex/bins"
	"struct-ex/file"
)

type Storage interface {
	SaveBinList(*bins.BinList, string) error
	ReadBinList(string) (*bins.BinList, error)
}

func SaveBinList(list *bins.BinList, filename string) error {
	data, err := json.Marshal(list)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func ReadBinList(filename string) (*bins.BinList, error) {
	content, err := file.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	list := &bins.BinList{}
	err = json.Unmarshal([]byte(content), list)
	if err != nil {
		return nil, err
	}
	return list, nil
}
