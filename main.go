package main

import (
	"fmt"
	"log"
	"struct-ex/api"
	"struct-ex/config"
	"struct-ex/file"
	"struct-ex/storage"

	"github.com/joho/godotenv"
)

func main() {
	fileService := file.NewFileService()

	storage := storage.NewStorageService(fileService)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := config.NewConfig()

	api := api.NewApi(storage, conf)
	fmt.Println(api)
}
