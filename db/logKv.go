package db

import (
	"log"
	"os"
)

type LogKvDB struct {
	ActivityFile *os.File          //	正在使用的File
	Indexes      map[string]uint32 //	维护的内存kv键值对
	DbPath       string            //	数据库文件存储目录
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
		Indexes:      map[string]uint32{},
		DbPath:       dbPath,
	}
	loadIndexFromDb(kvDb)
	return kvDb
}

//	加载索引从文件中
func loadIndexFromDb(kvDb *LogKvDB) {
	var offset int64=0
	buf:=make([]byte,10)
	:=kvDb.ActivityFile.ReadAt(buf,offset)
}