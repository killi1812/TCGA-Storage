package storage

import "os"

type Storage interface {
	Upload(file os.File) error
	Download(name string) (*os.File, error)
}
