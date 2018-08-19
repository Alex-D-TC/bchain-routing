package util

import (
	"bytes"
	"encoding/gob"
)

func MakeGobEncoder() (*gob.Encoder, *bytes.Buffer) {
	buffer := bytes.NewBuffer([]byte{})
	return gob.NewEncoder(buffer), buffer
}

func MakeGobDecoder() (*gob.Decoder, *bytes.Buffer) {
	buffer := bytes.NewBuffer([]byte{})
	return gob.NewDecoder(buffer), buffer
}

func GobEncode(data interface{}) ([]byte, error) {
	encoder, buffer := MakeGobEncoder()
	err := encoder.Encode(data)
	return buffer.Bytes(), err
}

func GobDecode(rawData []byte, destination interface{}) error {
	decoder, buffer := MakeGobDecoder()
	buffer.Write(rawData)
	return decoder.Decode(destination)
}
