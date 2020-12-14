package serialization

import "github.com/vmihailenco/msgpack/v5"

// MsgPackSerialization Serialization to Msg Pack Format
type MsgPackSerialization struct{}

// Decode Data Followings MsgPack
func (gobs MsgPackSerialization) Decode(data interface{}, obj interface{}) error {
	b, _ := data.([]byte)
	err := msgpack.Unmarshal(b, &obj)
	return err
}

// Encode Data Followings MsgPack
func (gobs MsgPackSerialization) Encode(data interface{}) interface{} {
	b, _ := msgpack.Marshal(&data)
	return b
}
