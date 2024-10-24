package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, _ := leveldb.OpenFile("./leveldb/data", nil)

	const count = 100000
	for i := 0; i < count; i++ {
		db.Put([]byte(fmt.Sprintf("key:%d", i)), []byte(fmt.Sprintf("value:%d", i)), nil)
	}

	for i := 0; i < count; i++ {
		value, _ := db.Get([]byte(fmt.Sprintf("key:%d", i)), nil)
		fmt.Println(string(value))
	}
	db.Close()
}
