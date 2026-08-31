package main

import (
	"fmt"
	"struct-ex/api"
	"struct-ex/file"
	"struct-ex/storage"
)

func main() {
	fileService := file.NewFileService()

	storage := storage.NewStorageService(fileService)

	api := api.NewApi(storage)
	fmt.Println(api)
}
