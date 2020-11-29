package inmemdb

import "sync"

// DbOperations that can be executed in the database
type DbOperations interface {
	Set(key interface{}, data interface{})
	Get(key interface{})
	Delete(key interface{})
}

// MemDb is a Volatile Memory Database
type MemDb struct {
	Db sync.Map
}

// NewMemDb Creates a new instance of MemDb
func NewMemDb() *MemDb {
	memDb := MemDb{Db: sync.Map{}}
	return &memDb
}

// Set an Object in the MemDb
func (mdb *MemDb) Set(key interface{}, data interface{}) (bool, error) {

	if _, ok := mdb.Db.Load(key); ok {
		return false, DuplicateKeyError{key: key}
	}

	mdb.Db.Store(key, data)

	return true, nil
}

// Get an Object in the MemDb
func (mdb *MemDb) Get(key interface{}) interface{} {
	value, _ := mdb.Db.Load(key)
	return value
}

// Delete an Object in the MemDb
func (mdb *MemDb) Delete(key interface{}) {
	_, ok := mdb.Db.Load(key)

	if ok {
		mdb.Db.Delete(key)
	}
}
