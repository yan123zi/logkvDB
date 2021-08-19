package main

import (
	"fmt"
	"logkvDB/db"
)

func main() {
	kvDb := db.OpenLogKvDb("tmp/kvDB")
	err := kvDb.Put("yan18", "zijiang")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	//buf:=make([]byte,2048)
	//binary.BigEndian.PutUint32(buf[0:4],10)
	//fmt.Println(buf)
	//fmt.Println(binary.BigEndian.Uint32(buf[0:4]))
	//buff:=make([]byte,2048)
	//file,err:=os.Open("tmp/kvDB/kvdb.data")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//redNum,err:=file.ReadAt(buff,0)
	//fe:=buf[3]
	//fmt.Println(fe)
	//fmt.Println(redNum,binary.BigEndian.Uint32(buff[22:26]),buff)
}
