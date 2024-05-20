package main

import "github.com/dgraph-io/badger"

func Handle(err error) {
	if err != nil {
		panic(err)
	}
}

func GetFromDB(db *badger.DB, key []byte) ([]byte, error) {
	var data []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			data = val
			return nil
		})
		return err
	})
	return data, err
}
func SetToDB(db *badger.DB, key []byte, value []byte) error {
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, value)
		return err
	})
	return err
}

func SetLastHashAndBlock(db *badger.DB, hash_key []byte, hash []byte, block_key []byte, block []byte) error {
	err := db.Update(func(txn *badger.Txn) error {
		err := txn.Set(hash_key, hash)
		if err != nil {
			return err
		}
		err = txn.Set(block_key, block)
		return err
	})
	return err
}
