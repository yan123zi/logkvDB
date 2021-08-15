package db

import (
	"io"
	"log"
	"os"
)

type LogKvDB struct {
	ActivityFile *DBfile          //	正在使用的File
	Indexes      map[string]int64 //	维护的内存kv键值对
	DbPath       string           //	数据库文件存储目录
}

func OpenLogKvDb(dbPath string) *LogKvDB {
	//	判断数据库目录是否存在,没有则创建
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		if err = os.MkdirAll(dbPath, os.ModePerm); err != nil {
			log.Fatalf("create dbPath err:%s", err.Error())
			return nil
		}
	}
	dbFile, err := NewDbFile(dbPath)
	if err != nil {
		log.Printf("newdbFile err:%s", err.Error())
		return nil
	}

	kvDb := &LogKvDB{
		ActivityFile: dbFile,
		Indexes:      map[string]int64{},
		DbPath:       dbPath,
	}
	loadIndexFromDb(kvDb)
	return kvDb
}

//	加载索引从文件中
func loadIndexFromDb(kvDb *LogKvDB) {
	var offset int64 = 0
	for {
		entry, err := kvDb.ActivityFile.Read(offset)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Printf("readAt err:%s", err.Error())
				return
			}
		}

		if _, ok := kvDb.Indexes[string(entry.Key)]; !ok && entry.GetLen() > 0 {
			kvDb.Indexes[string(entry.Key)] = entry.GetLen()
			offset += entry.GetLen()
		}
	}
	kvDb.ActivityFile.Offset = offset
}

func (db *LogKvDB) Put(key string, value string) error {
	entry := NewEntry(key, value, PUT)
	err := db.ActivityFile.Add(entry)
	if err != nil {
		return err
	}
	db.Indexes[key] = db.ActivityFile.Offset
	return nil
}
