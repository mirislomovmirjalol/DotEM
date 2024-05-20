package store

import (
	"go.etcd.io/bbolt"
	"log"
)

type BoltDBStore struct {
	DB *bbolt.DB
}

func GetStore() *BoltDBStore {
	db, err := bbolt.Open("/Users/user/Projects/go/DotEM/DotEM.db", 0600, nil)
	if err != nil {
		log.Fatal("Failed to open database: ", err)
	}
	return &BoltDBStore{DB: db}
}

func (store *BoltDBStore) Close() {
	if err := store.DB.Close(); err != nil {
		log.Fatal("Failed to close database: ", err)
	}
}

func (store *BoltDBStore) Set(bucketName, key string, value []byte) error {
	return store.DB.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return bucket.Put([]byte(key), value)
	})
}

func (store *BoltDBStore) Get(bucketName, key string) ([]byte, error) {
	var value []byte
	err := store.DB.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return bbolt.ErrBucketNotFound
		}
		bucketValue := bucket.Get([]byte(key))
		if bucketValue != nil {
			value = make([]byte, len(bucketValue))
			copy(value, bucketValue)
		}
		return nil
	})
	return value, err
}

func (store *BoltDBStore) Delete(bucketName, key string) error {
	return store.DB.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return bbolt.ErrBucketNotFound
		}
		return bucket.Delete([]byte(key))
	})
}

func (store *BoltDBStore) GetKeys(bucketName string) ([]string, error) {
	var keys []string
	err := store.DB.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(bucketName))
		if bucket == nil {
			return bbolt.ErrBucketNotFound
		}
		return bucket.ForEach(func(k, v []byte) error {
			keys = append(keys, string(k))
			return nil
		})
	})
	return keys, err
}
