package util

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"secondbit.org/wendy"
)

func NodeIDFromBytesSHA(entropySource []byte) wendy.NodeID {
	result := sha256.Sum256(entropySource)
	var id wendy.NodeID

	id[0] = binary.BigEndian.Uint64(result[:8])
	id[1] = binary.BigEndian.Uint64(result[8:16])

	return id
}

func NodeIDFromStringSHA(str string) wendy.NodeID {
	return NodeIDFromBytesSHA([]byte(str))
}

func NodeIDFromHexForm(str string) (wendy.NodeID, error) {
	bytes, err := hex.DecodeString(str)
	if err != nil {
		fmt.Println(err)
		return wendy.NodeID{}, err
	}

	id := wendy.NodeID{}
	id[0] = BytesToUint64(bytes[0:8])
	id[1] = BytesToUint64(bytes[8:])

	return id, nil
}

func NodeIDToString(id wendy.NodeID) string {
	bytes := append(Uint64ToBytes(id[0]), Uint64ToBytes(id[1])...)
	return hex.EncodeToString(bytes)
}
