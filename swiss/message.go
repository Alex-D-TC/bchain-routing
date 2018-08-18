package swiss

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/alex-d-tc/bchain-routing/util"
	"secondbit.org/wendy"
)

var messageHashFunction = crypto.SHA256
var messageHashFunctionImpl = sha256.New()

type Message struct {
	Sender   *wendy.NodeID
	Receiver *wendy.NodeID

	SenderPubKey *rsa.PublicKey

	RelayChain []relayBlock
	Payload    []byte

	Signature []byte
}

func MakeMessage(sender *wendy.NodeID, senderPubKey *rsa.PublicKey, senderPrivateKey *rsa.PrivateKey, receiver *wendy.NodeID, payload []byte) (*Message, error) {
	msg := Message{
		Sender:       sender,
		Receiver:     receiver,
		SenderPubKey: senderPubKey,
		RelayChain:   nil,
		Payload:      payload,
		Signature:    nil,
	}

	bytes, err := util.GobEncode(msg)
	if err != nil {
		return nil, err
	}

	sign, err := rsa.SignPSS(nil, senderPrivateKey, messageHashFunction, messageHashFunctionImpl.Sum(bytes), nil)
	if err != nil {
		return nil, err
	}

	msg.Signature = sign

	return &msg, err
}

func (msg *Message) Relay(id *wendy.NodeID, nextID *wendy.NodeID, senderPublicKey *rsa.PublicKey, senderPrivateKey *rsa.PrivateKey) error {

	currentID := msg.Sender
	var prevBlock *relayBlock
	if len(msg.RelayChain) > 0 {
		currentID = msg.RelayChain[0].NextID
		prevBlock = &msg.RelayChain[0]
	}

	if currentID != id {
		return errors.New("The supplied ID does not match the NextID of the latest block in the Proof of Relay chain")
	}

	block, err := makeRelayBlock(id, senderPublicKey, senderPrivateKey, nextID, prevBlock)
	if err != nil {
		return err
	}

	msg.RelayChain = append(msg.RelayChain, *block)
	return nil
}

func DefaultMessageProcessor(msg *Message) {
	fmt.Println("Processing message with the default processor...")
	fmt.Println(string(msg.Payload))
}
