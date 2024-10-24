package main

import (
	"fmt"

	"github.com/cockroachdb/pebble"
)

func main() {
	db, _ := pebble.Open("./data", &pebble.Options{})

	const count = 100000
	batch := db.NewBatch()
	for i := 0; i < count; i++ {
		key := fmt.Sprintf("key:%d", i)
		value := fmt.Sprintf("value:%d", i)
		batch.Set([]byte(key), []byte(value), pebble.NoSync)
	}
	batch.Commit(pebble.NoSync)

	db.Close()
}

// for i := 0; i < count; i++ {
// 	value, closer, _ := db.Get([]byte(fmt.Sprintf("key:%d", i)))
// 	fmt.Println(string(value))
// 	defer closer.Close()
// }
