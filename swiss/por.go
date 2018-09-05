package swiss

import (
	"crypto/ecdsa"

	"github.com/alex-d-tc/bchain-routing/net"
	"github.com/alex-d-tc/bchain-routing/util"
	"github.com/ethereum/go-ethereum/crypto"
)

type RelayBlock struct {
	ID     net.NodeID
	NextID net.NodeID
	PrevID net.NodeID

	PubKeyRaw     []byte
	PrevPubKeyRaw []byte

	Signature     []byte
	PrevSignature []byte
}

type validationRelayBlock struct {
	ID     net.NodeID
	NextID net.NodeID
	PrevID net.NodeID

	PubKeyRaw     []byte
	PrevPubKeyRaw []byte

	PrevSignature []byte
}

func makeRelayBlock(id net.NodeID, privKey *ecdsa.PrivateKey, nextID net.NodeID, prevRelayBlock *RelayBlock) (*RelayBlock, error) {
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

	hash, err := block.BlockHash256()
	if err != nil {
		return nil, err
	}

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

func (block *RelayBlock) BlockHash256() ([]byte, error) {

	encoded, err := util.GobEncode(*makeValidationRelayBlock(block))
	if err != nil {
		return nil, err
	}

	return crypto.Keccak256(encoded), nil
}

func (block *RelayBlock) ValidationBytes() ([]byte, error) {
	return util.GobEncode(*makeValidationRelayBlock(block))
}
