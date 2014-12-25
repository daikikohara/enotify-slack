package main_test

import (
	"io/ioutil"
	"os"
	"testing"

	. "github.com/daikikohara/enotify-slack"
	"github.com/stretchr/testify/assert"
)

const (
	dbfile   = "bolt.db"
	bucket   = "bucket"
	key      = "key"
	value    = "value"
	nonexist = "/path/to/nonexist"
	empty    = ""
)

func TestNewBolt(t *testing.T) {
	// case1: Success
	file, err := ioutil.TempFile(os.TempDir(), dbfile)
	if err != nil {
		t.Error(err.Error())
	}
	defer os.Remove(file.Name())
	bolt := NewBolt(file.Name(), bucket)
	defer bolt.Db.Close()

	assert := assert.New(t)
	assert.Equal(file.Name(), bolt.Dbfile, "boltdb file name is not correct")
	assert.Equal(bucket, bolt.Bucket, "boltdb bucket name is not correct")

	// case2: invalid path
	fn := func(file, bucket string) {
		defer func() {
			if r := recover(); r == nil {
				t.Fail()
			}
		}()
		NewBolt(file, bucket)
	}
	fn(nonexist, bucket)
}

func TestPut(t *testing.T) {
	// case1: bucket name is empty
	file, err := ioutil.TempFile(os.TempDir(), dbfile)
	if err != nil {
		t.Error(err.Error())
	}
	defer os.Remove(file.Name())
	bolt := NewBolt(file.Name(), empty)
	defer bolt.Db.Close()

	fn := func(k, v string, b *Bolt) {
		defer func() {
			if r := recover(); r == nil {
				t.Fail()
			}
		}()
		b.Put([]byte(k), []byte(v))
	}
	fn(key, value, bolt)

	// case2: key is empty
	file2, err := ioutil.TempFile(os.TempDir(), dbfile+"2")
	if err != nil {
		t.Error(err.Error())
	}
	defer os.Remove(file2.Name())
	bolt2 := NewBolt(file2.Name(), bucket)
	defer bolt2.Db.Close()
	fn(empty, value, bolt2)

	// case3: success
	// success case can be tested in TestExists
}

func TestExists(t *testing.T) {
	// case1:
	file, err := ioutil.TempFile(os.TempDir(), dbfile)
	if err != nil {
		t.Error(err.Error())
	}
	defer os.Remove(file.Name())
	bolt := NewBolt(file.Name(), bucket)
	defer bolt.Db.Close()

	assert := assert.New(t)
	assert.False(bolt.Exists([]byte(key)), "There has to be NO bucket in boltdb")
	bolt.Put([]byte(key), []byte(value))
	assert.False(bolt.Exists([]byte(nonexist)), "There has to be NO key in boltdb")
	assert.True(bolt.Exists([]byte(key)), "There has to be NO key in boltdb")

}
