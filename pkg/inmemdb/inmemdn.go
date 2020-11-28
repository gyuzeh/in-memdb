package inmemdb

// DbOperations that can be executed in the database
type DbOperations interface {
	Set(key string, data interface{})
	Get(key string)
	Delete(key string)
}

// MemDb is a Volatile Memory Database
type MemDb struct {
	Db map[string]interface{}
}

// NewMemDb Creates a new instance of MemDb
func NewMemDb() *MemDb {
	memDb := MemDb{Db: make(map[string]interface{})}
	return &memDb
}

// Set an Object in the MemDb
func (mdb *MemDb) Set(key string, data interface{}) (bool, error) {

	if _, ok := mdb.Db[key]; ok {
		return false, DuplicateKeyError{key: key}
	}

	mdb.Db[key] = data

	return true, nil
}

// Get an Object in the MemDb
func (mdb *MemDb) Get(key string) interface{} {
	return mdb.Db[key]
}

// Delete an Object in the MemDb
func (mdb *MemDb) Delete(key string) {
	_, ok := mdb.Db[key]

	if ok {
		delete(mdb.Db, key)
	}
}
