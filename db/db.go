package db

import (
	"errors"
	"io"
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
	entryByt := EncodeEntry(entry)
	if len(entryByt) == 0 {
		return errors.New("entry is null")
	}
	entryByt=append(entryByt,'\n')
	_, err := db.File.WriteAt(entryByt, db.Offset)
	if err != nil {
		log.Printf("Add entry to DbFile err:%s", err.Error())
		return err
	}
	db.Offset += entry.GetLen()
	return nil
}

func (db *DBfile) Read(offset int64) (*Entry, error) {
	buf := make([]byte, 2048)
	readNum, err := db.File.ReadAt(buf, offset)
	if err == io.EOF {
		if readNum == 0 {
			return nil, err
		}
	} else {
		return nil, err
	}

	entry := DecodeEntry(buf[:readNum])
	if entry.KeySize > 0 {
		entry.Key = buf[CommonFileLength : CommonFileLength+entry.KeySize]
		offset += int64(CommonFileLength + entry.KeySize)
	}
	if entry.ValueSize > 0 {
		entry.Value = buf[offset:]
		offset += int64(entry.ValueSize)
	}
	return entry, nil
}
