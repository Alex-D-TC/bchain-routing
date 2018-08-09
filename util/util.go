package util

import (
	"crypto/sha256"
	"encoding/binary"

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
