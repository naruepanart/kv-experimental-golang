package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	db, _ := leveldb.OpenFile("./data", nil)
	defer db.Close()

	const count = 100000
	batch := new(leveldb.Batch)
	for i := 0; i < count; i++ {
		key := fmt.Sprintf("key:%d", i)
		value := fmt.Sprintf("value:%d", i)
		batch.Put([]byte(key), []byte(value))
	}
	db.Write(batch, nil)
}

// for i := 0; i < count; i++ {
// 	value, _ := db.Get([]byte(fmt.Sprintf("key:%d", i)), nil)
// 	fmt.Println(string(value))
// }
