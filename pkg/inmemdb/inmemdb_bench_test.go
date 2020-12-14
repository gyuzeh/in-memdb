package inmemdb_test

import (
	"testing"

	"github.com/gyuzeh/in-memdb/pkg/inmemdb"
	"github.com/gyuzeh/in-memdb/pkg/serialization"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Colors      []string
	Size        []string
}

func BenchmarkInMemDbSetWithDefaultSerialization(b *testing.B) {
	// Arrange
	data := Product{
		ID:          12355634,
		Name:        "Shirt",
		Description: "This is my shirt and I love it! Buy me",
		Price:       300.2,
		Colors:      []string{"red", "yello", "blue"},
		Size:        []string{"S", "M", "L"},
	}

	memdb := inmemdb.New(serialization.PlainSerialization{})

	//Act
	for n := 0; n < b.N; n++ {
		memdb.Set(n, data)
	}
}

func BenchmarkInMemDbSetWithGobsSerialization(b *testing.B) {
	// Arrange
	data := Product{
		ID:          12355634,
		Name:        "Shirt",
		Description: "This is my shirt and I love it! Buy me",
		Price:       300.2,
		Colors:      []string{"red", "yello", "blue"},
		Size:        []string{"S", "M", "L"},
	}
	memdb := inmemdb.New(serialization.GobsSerialization{})

	//Act
	for n := 0; n < b.N; n++ {
		memdb.Set(n, data)
	}
}

func BenchmarkInMemDbSetWithMsgPackSerialization(b *testing.B) {
	// Arrange
	data := Product{
		ID:          12355634,
		Name:        "Shirt",
		Description: "This is my shirt and I love it! Buy me",
		Price:       300.2,
		Colors:      []string{"red", "yello", "blue"},
		Size:        []string{"S", "M", "L"},
	}
	memdb := inmemdb.New(serialization.MsgPackSerialization{})

	//Act
	for n := 0; n < b.N; n++ {
		memdb.Set(n, data)
	}
}

func BenchmarkInMemDbGetWithDefaultSerialization(b *testing.B) {
	// Arrange
	data := Product{
		ID:          12355634,
		Name:        "Shirt",
		Description: "This is my shirt and I love it! Buy me",
		Price:       300.2,
		Colors:      []string{"red", "yello", "blue"},
		Size:        []string{"S", "M", "L"},
	}
	memdb := inmemdb.New(serialization.PlainSerialization{})

	for n := 0; n < b.N; n++ {
		memdb.Set(n, data)
	}

	b.ResetTimer()

	// Act
	for n := 0; n < b.N; n++ {
		memdb.Get(n, data)
	}

}

func BenchmarkInMemDbGetWithGobsSerialization(b *testing.B) {
	// Arrange
	data := Product{
		ID:          12355634,
		Name:        "Shirt",
		Description: "This is my shirt and I love it! Buy me",
		Price:       300.2,
		Colors:      []string{"red", "yello", "blue"},
		Size:        []string{"S", "M", "L"},
	}
	memdb := inmemdb.New(serialization.GobsSerialization{})

	for n := 0; n < b.N; n++ {
		memdb.Set(n, data)
	}

	b.ResetTimer()

	// Act
	for n := 0; n < b.N; n++ {
		memdb.Get(n, data)
	}
}

func BenchmarkInMemDbGetWithMsgPackSerialization(b *testing.B) {
	// Arrange
	data := Product{
		ID:          12355634,
		Name:        "Shirt",
		Description: "This is my shirt and I love it! Buy me",
		Price:       300.2,
		Colors:      []string{"red", "yello", "blue"},
		Size:        []string{"S", "M", "L"},
	}
	memdb := inmemdb.New(serialization.MsgPackSerialization{})

	for n := 0; n < b.N; n++ {
		memdb.Set(n, data)
	}

	b.ResetTimer()

	// Act
	for n := 0; n < b.N; n++ {
		memdb.Get(n, data)
	}
}
