package serialization

import (
	"bytes"
	"encoding/gob"
)

// GobsSerialization Gobs Serialization
type GobsSerialization struct{}

// Decode Data Followings Gobs
func (gobs GobsSerialization) Decode(data interface{}, obj interface{}) error {
	b, _ := data.([]byte)
	d := gob.NewDecoder(bytes.NewReader(b))
	return d.Decode(obj)
}

// Encode Data Followings Gobs
func (gobs GobsSerialization) Encode(data interface{}) interface{} {
	var buf bytes.Buffer
	gob.NewEncoder(&buf).Encode(data)
	return buf.Bytes()
}
