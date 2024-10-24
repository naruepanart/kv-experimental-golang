package main

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
)

func main() {
	opt := badger.DefaultOptions("./data")
	opt.NumVersionsToKeep = 1
	opt.SyncWrites = false
	db, _ := badger.Open(opt)
	defer db.Close()

	const count = 100000
	wb := db.NewWriteBatch()
	defer wb.Cancel()
	for i := 0; i < count; i++ {
		key := fmt.Sprintf("key:%d", i)
		value := fmt.Sprintf("value:%d", i)
		wb.Set([]byte(key), []byte(value))
	}
	wb.Flush()
}

// for i := 0; i < count; i++ {
// 	db.View(func(txn *badger.Txn) error {
// 		item, _ := txn.Get([]byte(fmt.Sprintf("key:%d", i)))
// 		value, _ := item.ValueCopy(nil)
// 		fmt.Println(string(value))
// 		return nil
// 	})
// }
