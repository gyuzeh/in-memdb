package inmemdb

import (
	"testing"

	"github.com/gyuzeh/in-memdb/internals/serialization"
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
	memdb := NewMemDb(serialization.NullSerialization{})

	// Act
	err := memdb.Set(key, data)

	// Assert
	assert.Nil(t, err)
}

func TestSetWhenKeyAndObjectValueAreProvidedThenIsSaved(t *testing.T) {
	// Arrange
	key := "empKey"
	data := Employee{FirstName: "Rocky", LastName: "Sting", City: "London"}
	memdb := NewMemDb(serialization.NullSerialization{})

	// Act
	err := memdb.Set(key, data)

	// Assert
	assert.Nil(t, err)
}

func TestSetWhenValueWithExistingKeyIsProvidedThenItShouldFail(t *testing.T) {
	// Arrange
	key := "12345"
	data := Employee{FirstName: "Rocky", LastName: "Sting", City: "London"}
	memdb := NewMemDb(serialization.NullSerialization{})
	memdb.Set(key, data)

	// Act
	err := memdb.Set(key, data)

	// Assert
	assert.NotNil(t, err)
	assert.Equal(t, "Duplicated Key 12345", err.Error())
}

func TestGetWhenKeyExitsThenReturnValue(t *testing.T) {
	// Arrange
	key := "key"
	data := Employee{FirstName: "Rocky", LastName: "Sting", City: "London"}
	memdb := NewMemDb(serialization.NullSerialization{})
	memdb.Set(key, data)
	var employee Employee

	// Act
	err := memdb.Get(key, &employee)

	// Assert
	assert.Nil(t, err)
}

func TestGetWhenKeyExitsThenDelete(t *testing.T) {
	// Arrange
	key := "qwerty"
	memdb := NewMemDb(serialization.NullSerialization{})
	memdb.Set(key, "random value")
	var result Employee

	// Act
	memdb.Delete(key)

	// Assert
	memdb.Get(key, &result)
	assert.Equal(t, Employee{}, result)
}

func TestDeleteWhenKeyDoesNotExitsThenDontThrowError(t *testing.T) {
	// Arrange
	key := "unExistingKey"
	memdb := NewMemDb(serialization.NullSerialization{})

	// Act & Assert
	memdb.Delete(key)
}

func TestSetWhenSerializationMsgPackIsConfiguredThenSetAndGet(t *testing.T) {
	// Arrange
	key := "empKey"
	data := Employee{FirstName: "Rocky", LastName: "Sting", City: "London"}
	memdb := NewMemDb(serialization.MsgPackSerialization{})

	// Act
	err := memdb.Set(key, data)

	// Assert
	assert.Nil(t, err)
	var employee Employee
	memdb.Get(key, &employee)
	assert.Equal(t, employee, data)
}

func TestSetWhenSerializationGobsIsConfiguredThenSetAndGet(t *testing.T) {
	// Arrange
	key := "empKey"
	data := Employee{FirstName: "Rocky", LastName: "Sting", City: "London"}
	memdb := NewMemDb(serialization.GobsSerialization{})

	// Act
	err := memdb.Set(key, data)

	// Assert
	assert.Nil(t, err)
	var employee Employee
	memdb.Get(key, &employee)
	assert.Equal(t, employee, data)
}
