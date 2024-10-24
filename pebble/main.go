package main

import (
	"fmt"

	"github.com/cockroachdb/pebble"
)

func main() {
	db, _ := pebble.Open("./pebble/data", &pebble.Options{})

	const count = 100000
	batch := db.NewBatch()
	for i := 0; i < count; i++ {
		batch.Set([]byte(fmt.Sprintf("key:%d", i)), []byte(fmt.Sprintf("value:%d", i)), pebble.NoSync)
	}
	batch.Commit(pebble.NoSync)

	db.Close()
}

// for i := 0; i < count; i++ {
// 	value, closer, _ := db.Get([]byte(fmt.Sprintf("key:%d", i)))
// 	fmt.Println(string(value))
// 	defer closer.Close()
// }
