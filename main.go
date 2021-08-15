package main

import (
	"fmt"
	"logkvDB/db"
)

func main() {
	kvDb := db.OpenLogKvDb("tmp/kvDB")
	err := kvDb.Put("yan7", "zijiang")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
}
