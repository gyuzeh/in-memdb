package serialization

// Serialization Interface for Data
type Serialization interface {
	Decode(data interface{}, obj interface{}) error
	Encode(data interface{}) interface{}
}
