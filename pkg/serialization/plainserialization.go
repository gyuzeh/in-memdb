package serialization

// PlainSerialization does no serialization (plain objects)
type PlainSerialization struct{}

// Decode Data Followings Plain
func (plain PlainSerialization) Decode(data interface{}, obj interface{}) error {
	obj = data
	return nil
}

// Encode Data Followings Plain
func (plain PlainSerialization) Encode(data interface{}) interface{} {
	return data
}
