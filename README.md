# In Memory Database

Golang InMemory Database allows to abstract from your database implementation to help you in development by replacing the atual database with a stubbed.

## Quickstart

    go get github.com/gyuzeh/in-memdb


```go
import "github.com/gyuzeh/in-memdb/pkg/serialization"
import "github.com/gyuzeh/in-memdb/pkg/inmemdb"

func InMemDbExample() {

    key := "12355634"
    data := Product{
		ID:          12355634,
		Name:        "Shirt",
		Description: "This is my shirt and I love it! Buy me",
		Price:       300.2,
		Colors:      []string{"red", "yello", "blue"},
		Size:        []string{"S", "M", "L"},
    }
    
    memdb := inmemdb.New(serialization.PlainSerialization{})
    memdb.Set(key, data)
    memdb.Get(key, data)
    memdb.Delete(key)
}
```

## Available Serialization for InMemory Db
```go
    memdb := inmemdb.New(serialization.PlainSerialization{}) // plain objects in memory
    memdb := inmemdb.New(serialization.MsgPackSerialization{}) // serializes to MsgPack in Memory
    memdb := inmemdb.New(serialization.GobSerialization{}) // serializes to Gob in Memory
```

## Building 

    make build

## Running the tests

    make unit-test

## Running benchmark tests

    make benchmark-test 
