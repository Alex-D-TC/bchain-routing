package swiss

import (
	"crypto/ecdsa"
	"crypto/sha256"

	"github.com/alex-d-tc/bchain-routing/util"
	"secondbit.org/wendy"
)

type RelayBlock struct {
	ID     wendy.NodeID
	NextID wendy.NodeID
	PrevID wendy.NodeID

	PubKey     ecdsa.PublicKey
	PrevPubKey ecdsa.PublicKey

	Signature     util.EcdsaSignature
	PrevSignature util.EcdsaSignature
}

type validationRelayBlock struct {
	PrevID        wendy.NodeID
	PrevPubKey    ecdsa.PublicKey
	PrevSignature util.EcdsaSignature

	ID     wendy.NodeID
	PubKey ecdsa.PublicKey

	NextID wendy.NodeID
}

func makeRelayBlock(id wendy.NodeID, privKey *ecdsa.PrivateKey, nextID wendy.NodeID, prevRelayBlock *RelayBlock) (*RelayBlock, error) {
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

	signature, err := util.Sign(privKey, hash)
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
	return util.JSONEncode(*makeValidationRelayBlock(block))
}
