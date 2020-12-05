package inmemdb

import (
	"sync"

	"github.com/gyuzeh/in-memdb/internals/encoding"
)

// Db that can be executed in the database
type Db interface {
	Set(key interface{}, data interface{}) error
	Get(key interface{}, data interface{}) error
	Delete(key interface{})
}

// MemDb is a Volatile Memory Database
type memDb struct {
	storage sync.Map
	enc     encoding.Encoding
}

// NewMemDb Creates a new instance of MemDb
func NewMemDb() Db {
	db := memDb{storage: sync.Map{}, enc: &encoding.NullConverter{}}
	return &db
}

// UseEnconding changes the default enconfing in MemDb
func (mdb *memDb) UseEnconding(enconding encoding.Encoding) {
	mdb.enc = enconding
}

// Set an Object in the MemDb
func (mdb *memDb) Set(key interface{}, data interface{}) error {
	if _, ok := mdb.storage.Load(key); ok {
		return DuplicateKeyError{key: key}
	}

	mdb.storage.Store(key, mdb.enc.Encode(data))

	return nil
}

// Get an Object in the MemDb
func (mdb *memDb) Get(key interface{}, data interface{}) error {
	value, _ := mdb.storage.Load(key)
	return mdb.enc.Decode(value, data)
}

// Delete an Object in the MemDb
func (mdb *memDb) Delete(key interface{}) {
	_, ok := mdb.storage.Load(key)

	if ok {
		mdb.storage.Delete(key)
	}
}
