package main

import (
	"fmt"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

func main() {
	db, err := leveldb.OpenFile("./data", nil)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	const count = 100_000

	batch := new(leveldb.Batch)
	for i := 0; i < count; i++ {
		key := fmt.Sprintf("key:%d", i)
		value := fmt.Sprintf("value:%d", i)
		batch.Put([]byte(key), []byte(value))
	}

	writeOptions := &opt.WriteOptions{
		Sync: false,
	}

	err = db.Write(batch, writeOptions)
	if err != nil {
		panic(err)
	}
}

// for i := 0; i < count; i++ {
// 	value, _ := db.Get([]byte(fmt.Sprintf("key:%d", i)), nil)
// 	fmt.Println(string(value))
// }
