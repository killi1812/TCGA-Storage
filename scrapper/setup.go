package scrapper

import (
	"TCGA-storage/config"
	"TCGA-storage/storage"
	"sync"
)

const prog = "PPPK_Scrapper"
const data = "Data"

var minioStorage *storage.MinioStorage

// TODO: add contex and abbyliti to cancle scraping
var fileLock sync.Mutex
var path string

func Setup() error {
	minioStorage = storage.New()
	fileLock = sync.Mutex{}

	path = config.Conf.ScrapperPath

	return nil
}
