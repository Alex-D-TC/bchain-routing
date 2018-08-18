package swiss

import (
	"bytes"
	"crypto"
	"crypto/sha256"
	"encoding/gob"
)

type relayBlock struct {
	pubKey        crypto.PublicKey
	prevPubKey    crypto.PublicKey
	prevHash      [sha256.Size]byte
	prevSignature []byte
}

type validationRelayBlock struct {
	pubKey        crypto.PublicKey
	prevPubKey    crypto.PublicKey
	prevHash      [sha256.Size]byte
	prevSignature []byte
}

func emptyRelayBlock(pubKey crypto.PublicKey) *relayBlock {
	return &relayBlock{
		prevHash: [sha256.Size]byte{},
	}
}

func makeRelayBlock(pubKey crypto.PublicKey, prevPubKey crypto.PublicKey, prevHash [sha256.Size]byte) *relayBlock {
	return &relayBlock{
		pubKey:     pubKey,
		prevPubKey: prevPubKey,
		prevHash:   prevHash,
	}
}

func makeValidationRelayBlock(block *relayBlock) *validationRelayBlock {
	return &validationRelayBlock{
		pubKey:        block.pubKey,
		prevPubKey:    block.prevPubKey,
		prevHash:      block.prevHash,
		prevSignature: block.prevSignature,
	}
}

func (block *relayBlock) ValidationBytes() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	encoder.Encode(*makeValidationRelayBlock(block))
	return result.Bytes()
}

func Validate(blocks []relayBlock) bool {
	blocksCount := len(blocks)

	if !bytes.Equal(blocks[blocksCount-1].prevHash[:], make([]byte, sha256.Size)) {
		return false
	}

	for i := blocksCount - 2; i >= 0; i-- {
		prevHash := sha256.Sum256(blocks[i+1].ValidationBytes())
		if !bytes.Equal(prevHash[:], blocks[i].prevHash[:]) {
			return false
		}
	}

	return true
}
