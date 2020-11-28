package inmemdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Employee struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	City      string `json:"city"`
}

func TestSetWhenKeyAndStringValueAreProvidedThenIsSaved(t *testing.T) {
	// Arrange
	key := "12345"
	data := `
    {
        "firstname": "Rocky",
        "lastname": "Sting",
        "city": "London"
    }`
	memdb := NewMemDb()

	// Act
	result, err := memdb.Set(key, data)

	// Assert
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestSetWhenKeyAndObjectValueAreProvidedThenIsSaved(t *testing.T) {
	// Arrange
	key := "empKey"
	data := Employee{FirstName: "Rocky", LastName: "Sting", City: "London"}
	memdb := NewMemDb()

	// Act
	result, err := memdb.Set(key, data)

	// Assert
	assert.Nil(t, err)
	assert.True(t, result)
}

func TestSetWhenValueWithExistingKeyIsProvidedThenItShouldFail(t *testing.T) {
	// Arrange
	key := "12345"
	data := Employee{FirstName: "Rocky", LastName: "Sting", City: "London"}
	memdb := NewMemDb()
	memdb.Set(key, data)

	// Act
	result, err := memdb.Set(key, data)

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, "Duplicated Key 12345", err.Error())
	assert.False(t, result)
}

func TestGetWhenKeyExitsThenReturnValue(t *testing.T) {
	// Arrange
	key := "key"
	data := Employee{FirstName: "Rocky", LastName: "Sting", City: "London"}
	memdb := NewMemDb()
	memdb.Set(key, data)

	// Act
	result := memdb.Get(key)

	// Assert
	assert.Equal(t, data, result)
}

func TestGetWhenKeyExitsThenDelete(t *testing.T) {
	// Arrange
	key := "qwerty"
	memdb := NewMemDb()
	memdb.Set(key, "random value")

	// Act
	memdb.Delete(key)

	// Assert
	assert.Nil(t, memdb.Get(key))
}

func TestDeleteWhenKeyDoesNotExitsThenDontThrowError(t *testing.T) {
	// Arrange
	key := "unExistingKey"
	memdb := NewMemDb()

	// Act & Assert
	memdb.Delete(key)
}
