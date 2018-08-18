package swiss

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"

	"github.com/alex-d-tc/bchain-routing/util"
	"secondbit.org/wendy"
)

var porHashFunc = crypto.SHA256
var porHashFuncImpl = sha256.New()

type relayBlock struct {
	ID     *wendy.NodeID
	NextID *wendy.NodeID
	PrevID *wendy.NodeID

	PubKey     *rsa.PublicKey
	PrevPubKey *rsa.PublicKey

	Signature     []byte
	PrevSignature []byte
}

type validationRelayBlock struct {
	PrevID        *wendy.NodeID
	PrevPubKey    *rsa.PublicKey
	PrevSignature []byte

	ID     *wendy.NodeID
	PubKey *rsa.PublicKey

	NextID *wendy.NodeID
}

func makeRelayBlock(id *wendy.NodeID, pubKey *rsa.PublicKey, privKey *rsa.PrivateKey, nextID *wendy.NodeID, prevRelayBlock *relayBlock) (*relayBlock, error) {
	block := relayBlock{
		PrevID:        prevRelayBlock.ID,
		PrevPubKey:    prevRelayBlock.PubKey,
		PrevSignature: prevRelayBlock.Signature,
		ID:            id,
		PubKey:        pubKey,
		NextID:        nextID,
	}

	blockBytes, err := block.ValidationBytes()
	if err != nil {
		return nil, err
	}

	signature, err := rsa.SignPSS(nil, privKey, porHashFunc, porHashFuncImpl.Sum(blockBytes), nil)
	if err == nil {
		block.Signature = signature
	}

	return &block, err
}

func makeValidationRelayBlock(block *relayBlock) *validationRelayBlock {
	return &validationRelayBlock{
		PrevID:        block.PrevID,
		PrevPubKey:    block.PrevPubKey,
		PrevSignature: block.PrevSignature,
		ID:            block.ID,
		PubKey:        block.PubKey,
		NextID:        block.NextID,
	}
}

func (block *relayBlock) ValidationBytes() ([]byte, error) {
	return util.GobEncode(*makeValidationRelayBlock(block))
}

func Validate(blocks []relayBlock) bool {
	return true
}
