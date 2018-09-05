package util

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"

	"github.com/alex-d-tc/bchain-routing/net"
)

func NodeIDFromBytesSHA(entropySource []byte) net.NodeID {
	result := sha256.Sum256(entropySource)
	var id net.NodeID

	id[0] = binary.BigEndian.Uint64(result[:8])
	id[1] = binary.BigEndian.Uint64(result[8:16])

	return id
}

func NodeIDFromStringSHA(str string) net.NodeID {
	return NodeIDFromBytesSHA([]byte(str))
}

func NodeIDFromHexForm(str string) (net.NodeID, error) {
	bytes, err := hex.DecodeString(str)
	if err != nil {
		return net.NodeID{}, err
	}

	id := net.NodeID{}
	id[0] = BytesToUint64(bytes[0:8])
	id[1] = BytesToUint64(bytes[8:])

	return id, nil
}

func NodeIDToString(id net.NodeID) string {
	bytes := append(Uint64ToBytes(id[0]), Uint64ToBytes(id[1])...)
	return hex.EncodeToString(bytes)
}
