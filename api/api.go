package api

import (
	"struct-ex/config"
	"struct-ex/storage"
)

type Api struct {
	Storage storage.Storage
	Conf    *config.Config
}

func NewApi(storage storage.Storage, conf *config.Config) Api {
	return Api{Storage: storage, Conf: conf}
}
