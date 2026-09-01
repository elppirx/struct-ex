package main

import (
	"flag"
	"fmt"
	"log"
	"struct-ex/api"
	"struct-ex/bins"
	"struct-ex/config"
	"struct-ex/file"
	"struct-ex/storage"

	"github.com/joho/godotenv"
)

const storageFile = "local_bins.json"

func main() {
	createFlag := flag.Bool("create", false, "создать бин")
	updateFlag := flag.Bool("update", false, "обновить бин")
	deleteFlag := flag.Bool("delete", false, "удалить бин")
	getFlag := flag.Bool("get", false, "получить бин")
	listFlag := flag.Bool("list", false, "список бинов")

	fileFlag := flag.String("file", "", "путь к файлу с содержимым")
	nameFlag := flag.String("name", "", "имя бина")
	idFlag := flag.String("id", "", "id бина")

	flag.Parse()

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := config.NewConfig()
	fileService := file.NewFileService()
	storageService := storage.NewStorageService(fileService)
	apiService := api.NewApi(storageService, conf)

	switch {
	case *createFlag:
		handleCreate(apiService, storageService, fileService, *fileFlag, *nameFlag)
	case *updateFlag:
		handleUpdate(apiService, fileService, *fileFlag, *idFlag)
	case *deleteFlag:
		handleDelete(apiService, storageService, *idFlag)
	case *getFlag:
		handleGet(apiService, *idFlag)
	case *listFlag:
		handleList(storageService)
	default:
		fmt.Println("Не указана команда. Используйте --create, --update, --delete, --get или --list")
	}
}

func handleCreate(a api.Api, s storage.Storage, f file.File, filePath, name string) {
	if filePath == "" || name == "" {
		fmt.Println("нужны --file и --name")
		return
	}

	content, err := f.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	newBin, err := a.CreateBin(name, []byte(content))
	if err != nil {
		fmt.Println(err)
		return
	}

	list, err := s.ReadBinList(storageFile)
	if err != nil {
		list = bins.NewBinList([]bins.Bin{})
	}

	list.Bins = append(list.Bins, *newBin)

	if err := s.SaveBinList(list, storageFile); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Создан бин:", newBin.ID)
}

func handleUpdate(a api.Api, f file.File, filePath, id string) {
	if filePath == "" || id == "" {
		fmt.Println("нужны --file и --id")
		return
	}

	content, err := f.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := a.UpdateBin(id, []byte(content)); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Бин обновлён:", id)
}

func handleDelete(a api.Api, s storage.Storage, id string) {
	if id == "" {
		fmt.Println("нужен --id")
		return
	}

	if err := a.DeleteBin(id); err != nil {
		fmt.Println(err)
		return
	}

	list, err := s.ReadBinList(storageFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	newBins := make([]bins.Bin, 0)
	for _, b := range list.Bins {
		if b.ID != id {
			newBins = append(newBins, b)
		}
	}
	list.Bins = newBins

	if err := s.SaveBinList(list, storageFile); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Бин удалён:", id)
}

func handleGet(a api.Api, id string) {
	if id == "" {
		fmt.Println("нужен --id")
		return
	}

	content, err := a.GetBin(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(content)
}

func handleList(s storage.Storage) {
	list, err := s.ReadBinList(storageFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, b := range list.Bins {
		fmt.Println(b.ID, b.Name)
	}
}
