package main

import (
	"fmt"

	"log"

	bolt "go.etcd.io/bbolt"
)

func main() {
	db, err := bolt.Open("./data.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	const count = 100000
	db.Batch(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("mybucket"))
		if err != nil {
			return err
		}
		for i := 0; i < count; i++ {
			key := fmt.Sprintf("key:%d", i)
			value := fmt.Sprintf("value:%d", i)
			if err := bucket.Put([]byte(key), []byte(value)); err != nil {
				return err
			}
		}
		return nil
	})
}

// db.View(func(tx *bolt.Tx) error {
// 	bucket := tx.Bucket([]byte("mybucket"))
// 	if bucket == nil {
// 		return fmt.Errorf("bucket not found")
// 	}

// 	for i := 0; i < count; i++ {
// 		key := fmt.Sprintf("key:%d", i)
// 		value := bucket.Get([]byte(key))
// 		if value == nil {
// 			return fmt.Errorf("key %s not found", key)
// 		}
// 		fmt.Println(string(value))
// 	}
// 	return nil
// })
