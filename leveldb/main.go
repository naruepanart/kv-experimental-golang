package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, _ := leveldb.OpenFile("./leveldb/data", nil)
	defer db.Close()

	const count = 100000
	batch := new(leveldb.Batch)
	for i := 0; i < count; i++ {
		batch.Put([]byte(fmt.Sprintf("key:%d", i)), []byte(fmt.Sprintf("value:%d", i)))
	}
	db.Write(batch, nil)
}

// for i := 0; i < count; i++ {
// 	value, _ := db.Get([]byte(fmt.Sprintf("key:%d", i)), nil)
// 	fmt.Println(string(value))
// }
