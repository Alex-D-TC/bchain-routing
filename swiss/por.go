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

	PubKeyRaw     []byte
	PrevPubKeyRaw []byte

	Signature     []byte
	PrevSignature []byte
}

type validationRelayBlock struct {
	PrevID        wendy.NodeID
	PrevPubKeyRaw []byte
	PrevSignature []byte

	ID        wendy.NodeID
	PubKeyRaw []byte

	NextID wendy.NodeID
}

func makeRelayBlock(id wendy.NodeID, privKey *ecdsa.PrivateKey, nextID wendy.NodeID, prevRelayBlock *RelayBlock) (*RelayBlock, error) {
	block := RelayBlock{
		ID:        id,
		PubKeyRaw: util.MarshalPubKey(&privKey.PublicKey),
		NextID:    nextID,
	}

	if prevRelayBlock != nil {
		block.PrevID = prevRelayBlock.ID
		block.PrevPubKeyRaw = prevRelayBlock.PubKeyRaw
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
		PrevPubKeyRaw: block.PrevPubKeyRaw,
		PrevSignature: block.PrevSignature,
		ID:            block.ID,
		PubKeyRaw:     block.PubKeyRaw,
		NextID:        block.NextID,
	}
}

func (block *RelayBlock) ValidationBytes() ([]byte, error) {
	return util.JSONEncode(*makeValidationRelayBlock(block))
}
