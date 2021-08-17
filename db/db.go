package db

import (
	"bytes"
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
	//before:=[]byte{'\n'}
	entryByt := EncodeEntry(entry)
	if len(entryByt) == 0 {
		return errors.New("entry is null")
	}
	//before=append(before,entryByt...)
	//entryByt = append(entryByt, '\n')
	_, err := db.File.WriteAt(entryByt, db.Offset)
	if err != nil {
		log.Printf("Add entry to DbFile err:%s", err.Error())
		return err
	}
	db.Offset += entry.GetLen()
	return nil
}

func (db *DBfile) Read(offset int64) (*Entry, error) {
	buf := make([]byte, CommonFileLength)
	readNum, err := db.File.ReadAt(buf, offset)
	if err == io.EOF {
		if readNum == 0 || bytes.Equal(buf[:readNum], []byte{'\n'}) {
			return nil, err
		}
	} else if err!=nil {
		return nil, err
	}

	entry := DecodeEntry(buf[:readNum])
	if entry.KeySize > 0 {
		buf:=make([]byte,entry.KeySize)
		offset += int64(CommonFileLength)
		redNums,err:=db.File.ReadAt(buf,offset)
		if err != nil {
			return nil, err
		}
		entry.Key = buf[:redNums]

	}
	if entry.ValueSize > 0 {
		buf:=make([]byte,entry.ValueSize)
		offset += int64(CommonFileLength+entry.KeySize)
		redNums,err:=db.File.ReadAt(buf,offset)
		if err != nil &&err!=io.EOF{
			return nil, err
		}
		entry.Value = buf[:redNums]
	}
	return entry, nil
}
