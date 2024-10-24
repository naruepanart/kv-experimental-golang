package main

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
)

func main() {
	db, _ := badger.Open(badger.DefaultOptions("./badger/data"))
	defer db.Close()

	const count = 100000
	for i := 0; i < count; i++ {
		db.Update(func(txn *badger.Txn) error {
			return txn.Set([]byte(fmt.Sprintf("key:%d", i)), []byte(fmt.Sprintf("value:%d", i)))
		})
	}

	for i := 0; i < count; i++ {
		db.View(func(txn *badger.Txn) error {
			item, _ := txn.Get([]byte(fmt.Sprintf("key:%d", i)))
			value, _ := item.ValueCopy(nil)
			fmt.Println(string(value))
			return nil
		})
	}
}
