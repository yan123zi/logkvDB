package db

import (
	"log"
	"os"
)

const (
	ActivityDBFile = "kvdb.data"
	OldDBFile      = "old%d.kvdb.data"
)

type DBfile struct {
}

func NewDbFile(dbPath string) (*os.File, error) {
	fileName := ActivityDBFile + string(os.PathSeparator) + ActivityDBFile
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Printf("create kvdb err:%s", err.Error())
		return nil, err
	}
	return file, nil
}
