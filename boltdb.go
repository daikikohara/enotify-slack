package main

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

// Bolt holds a pointer of bolt.DB and 2 strings for dbfile and bucket name.
// boltdb is used to store event data in the following key-value format.
//  key - event-provider-name:event-id
//  value - detail of the event
type Bolt struct {
	Db     *bolt.DB
	Dbfile string
	Bucket string
}

// NewBolt opens boltdb to initialize Bolt.
// The path of the db file that stores event data is specified in the configuration file(yaml).
func NewBolt(dbfile, bucket string) *Bolt {
	db, err := bolt.Open(dbfile, 0644, nil)
	if err != nil {
		log.Panicln(err)
	}
	return &Bolt{db, dbfile, bucket}
}

// Exists checks if a key passed via the argument already exists in boltdb.
func (self *Bolt) Exists(key []byte) bool {
	err := self.Db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(self.Bucket))
		if bucket == nil {
			return nil
		}
		if bucket.Get(key) != nil {
			return fmt.Errorf("already registered")
		}
		return nil
	})
	if err != nil {
		return true
	}
	return false
}

// Put puts a key-value pair of an event data.
func (self *Bolt) Put(key, value []byte) {
	self.Db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte(self.Bucket))
		if err != nil {
			log.Panicln(err)
		}
		err = bucket.Put(key, value)
		if err != nil {
			log.Panicln(err)
		}
		return nil
	})
}
