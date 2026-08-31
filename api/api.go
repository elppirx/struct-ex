package api

import "struct-ex/storage"

type Api struct {
	Storage storage.Storage
}

func NewApi(storage storage.Storage) Api {
	return Api{Storage: storage}
}
