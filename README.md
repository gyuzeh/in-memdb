# In Memory Database

Golang InMemory database that allows the usage of several serialization protocols

## Quickstart

    go get github.com/gyuzeh/in-memdb


```go
import "github.com/gyuzeh/in-memdb/internals/serialization"
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
    
    memdb := inmemdb.New(serialization.NullSerialization{})
    memdb.Set(key, data)
    memdb.Get(key, data)
    memdb.Delete(key)
}
```

## Running the tests

    go test -v ./... 

## Running benchmark tests

    go test bench=. ./... 
