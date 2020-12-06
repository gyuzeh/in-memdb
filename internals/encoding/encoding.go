package encoding

import (
	"bytes"
	"encoding/gob"
)

// Encoding Interface for Data
type Encoding interface {
	Decode(data interface{}, obj interface{}) error
	Encode(data interface{}) interface{}
}

// GobsConverter Gobs Enconding
type GobsConverter struct{}

// Decode Data Followings Gobs
func (gobs GobsConverter) Decode(data interface{}, obj interface{}) error {
	b, _ := data.([]byte)
	d := gob.NewDecoder(bytes.NewReader(b))
	return d.Decode(obj)
}

// Encode Data Followings Gobs
func (gobs GobsConverter) Encode(data interface{}) interface{} {
	var buf bytes.Buffer
	gob.NewEncoder(&buf).Encode(data)
	return buf.Bytes()
}

// NullConverter does no enconding (plain objects)
type NullConverter struct{}

// Decode Data Followings Gobs
func (gobs NullConverter) Decode(data interface{}, obj interface{}) error {
	obj = data
	return nil
}

// Encode Data Followings Gobs
func (gobs NullConverter) Encode(data interface{}) interface{} {
	return data
}
