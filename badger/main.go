package main

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
)

func main() {
	opt := badger.DefaultOptions("./badger/data")
	opt.NumVersionsToKeep = 1
	opt.SyncWrites = false
	db, _ := badger.Open(opt)
	defer db.Close()

	const count = 100000
	wb := db.NewWriteBatch()
	defer wb.Cancel()
	for i := 0; i < count; i++ {
		wb.Set([]byte(fmt.Sprintf("key:%d", i)), []byte(fmt.Sprintf("value:%d", i)))
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
