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
	File   *os.File
	Offset int64
}

func NewDbFile(dbPath string) (*DBfile, error) {
	fileName := dbPath + string(os.PathSeparator) + ActivityDBFile
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Printf("create kvdb err:%s", err.Error())
		return nil, err
	}
	return &DBfile{File: file}, nil
}

func (db *DBfile) Add(entry *Entry) error {
	_, err := db.File.WriteAt(EncodeEntry(entry), db.Offset)
	if err != nil {
		log.Printf("Add entry to DbFile err:%s", err.Error())
		return err
	}
	db.Offset+=entry.GetLen()
	return nil
}
