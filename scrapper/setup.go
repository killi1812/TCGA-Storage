package scrapper

import (
	"TCGA-storage/storage"
	"sync"
)

const prog = "PPPK_Scrapper"
const path = ""
const data = "Data"

var minioStorage *storage.MinioStorage
var fileLock sync.Mutex

func Setup() error {
	sync.OnceFunc(func() {
		minioStorage = storage.New()
		fileLock = sync.Mutex{}
	})
	return nil
}
