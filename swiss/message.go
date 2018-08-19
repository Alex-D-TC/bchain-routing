package swiss

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"errors"
	"fmt"

	"github.com/alex-d-tc/bchain-routing/util"
	"secondbit.org/wendy"
)

type Message struct {
	Sender   wendy.NodeID
	Receiver wendy.NodeID

	SenderPubKey rsa.PublicKey

	RelayChain []RelayBlock
	Payload    []byte

	Signature []byte
}

func MakeMessage(sender wendy.NodeID, senderPrivateKey *rsa.PrivateKey, receiver wendy.NodeID, payload []byte) (*Message, error) {
	msg := Message{
		Sender:       sender,
		Receiver:     receiver,
		SenderPubKey: senderPrivateKey.PublicKey,
		RelayChain:   nil,
		Payload:      payload,
		Signature:    []byte{},
	}

	bytes, err := util.GobEncode(msg)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(bytes)

	sign, err := rsa.SignPSS(rand.Reader, senderPrivateKey, crypto.SHA256, hash[:], nil)
	if err != nil {
		return nil, err
	}

	msg.Signature = sign

	return &msg, err
}

func (msg *Message) Relay(id wendy.NodeID, nextID wendy.NodeID, senderPrivateKey *rsa.PrivateKey) error {

	currentID := msg.Sender
	var prevBlock *RelayBlock
	if len(msg.RelayChain) > 0 {
		currentID = msg.RelayChain[0].NextID
		prevBlock = &msg.RelayChain[0]
	}

	if currentID != id {
		return errors.New("The supplied ID does not match the NextID of the latest block in the Proof of Relay chain")
	}

	block, err := makeRelayBlock(id, senderPrivateKey, nextID, prevBlock)
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
