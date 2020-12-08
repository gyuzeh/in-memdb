package serialization

// NullSerialization does no serialization (plain objects)
type NullSerialization struct{}

// Decode Data Followings Gobs
func (gobs NullSerialization) Decode(data interface{}, obj interface{}) error {
	obj = data
	return nil
}

// Encode Data Followings Gobs
func (gobs NullSerialization) Encode(data interface{}) interface{} {
	return data
}
