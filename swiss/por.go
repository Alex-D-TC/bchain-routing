package swiss

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"

	"github.com/alex-d-tc/bchain-routing/util"
	"secondbit.org/wendy"
)

type RelayBlock struct {
	ID     wendy.NodeID
	NextID wendy.NodeID
	PrevID wendy.NodeID

	PubKey     rsa.PublicKey
	PrevPubKey rsa.PublicKey

	Signature     []byte
	PrevSignature []byte
}

type validationRelayBlock struct {
	PrevID        wendy.NodeID
	PrevPubKey    rsa.PublicKey
	PrevSignature []byte

	ID     wendy.NodeID
	PubKey rsa.PublicKey

	NextID wendy.NodeID
}

func makeRelayBlock(id wendy.NodeID, privKey *rsa.PrivateKey, nextID wendy.NodeID, prevRelayBlock *RelayBlock) (*RelayBlock, error) {
	block := RelayBlock{
		ID:     id,
		PubKey: privKey.PublicKey,
		NextID: nextID,
	}

	if prevRelayBlock != nil {
		block.PrevID = prevRelayBlock.ID
		block.PrevPubKey = prevRelayBlock.PubKey
		block.PrevSignature = prevRelayBlock.Signature
	}

	blockBytes, err := block.ValidationBytes()
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(blockBytes)

	signature, err := rsa.SignPSS(rand.Reader, privKey, crypto.SHA256, hash[:], nil)
	if err == nil {
		block.Signature = signature
	}

	return &block, err
}

func makeValidationRelayBlock(block *RelayBlock) *validationRelayBlock {
	return &validationRelayBlock{
		PrevID:        block.PrevID,
		PrevPubKey:    block.PrevPubKey,
		PrevSignature: block.PrevSignature,
		ID:            block.ID,
		PubKey:        block.PubKey,
		NextID:        block.NextID,
	}
}

func (block *RelayBlock) ValidationBytes() ([]byte, error) {
	return util.GobEncode(*makeValidationRelayBlock(block))
}

func Validate(blocks []RelayBlock) bool {
	return true
}
