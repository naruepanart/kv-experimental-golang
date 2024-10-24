package main

import (
	"fmt"

	"github.com/cockroachdb/pebble"
)

func main() {
	db, _ := pebble.Open("./data", &pebble.Options{})
	defer db.Close()

	const count = 100_000

	batch := db.NewBatch()
	defer batch.Close()

	for i := 0; i < count; i++ {
		key := fmt.Sprintf("key:%d", i)
		value := fmt.Sprintf("value:%d", i)
		batch.Set([]byte(key), []byte(value), nil)
	}

	batch.Commit(pebble.NoSync)
}

// for i := 0; i < count; i++ {
// 	value, closer, _ := db.Get([]byte(fmt.Sprintf("key:%d", i)))
// 	fmt.Println(string(value))
// 	defer closer.Close()
// }
