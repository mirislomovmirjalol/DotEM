package store

import (
	"go.etcd.io/bbolt"
)

type BoltDBStore struct {
	DB *bbolt.DB
}

func NewBoltDBStore(dbPath string) (*BoltDBStore, error) {
	db, err := bbolt.Open(dbPath, 0600, nil)
	if err != nil {
		return nil, err
	}
	return &BoltDBStore{DB: db}, nil
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
		value = bucket.Get([]byte(key))
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
