package swiss

import (
	"bytes"
	"crypto/ecdsa"
	"errors"
	"fmt"

	"github.com/alex-d-tc/bchain-routing/util"
	"github.com/ethereum/go-ethereum/crypto"
	"secondbit.org/wendy"
)

type Message struct {
	Sender   wendy.NodeID
	Receiver wendy.NodeID

	SenderPubKeyRaw []byte

	RelayChain []RelayBlock

	Payload []byte

	DataWithPayloadHash      []byte
	DataWithPayloadSignature []byte

	DataWithPayloadSizeHash      []byte
	DataWithPayloadSizeSignature []byte
}

func MakeMessage(sender wendy.NodeID, senderPrivateKey *ecdsa.PrivateKey, receiver wendy.NodeID, payload []byte) (*Message, error) {
	msg := Message{
		Sender:          sender,
		Receiver:        receiver,
		SenderPubKeyRaw: util.MarshalPubKey(&senderPrivateKey.PublicKey),
		RelayChain:      nil,
		Payload:         payload,
	}

	bytes, err := util.GobEncode(msg)
	if err != nil {
		return nil, err
	}

	hash := crypto.Keccak256(bytes)

	sign, err := util.Sign(senderPrivateKey, hash)
	if err != nil {
		return nil, err
	}

	msg.DataWithPayloadHash = hash
	msg.DataWithPayloadSignature = sign

	bytes, err = util.GobEncode(struct {
		Sender   wendy.NodeID
		Receiver wendy.NodeID

		SenderPubKeyRaw []byte

		RelayChain []RelayBlock

		PayloadSize uint64
	}{
		Sender:          msg.Sender,
		Receiver:        msg.Receiver,
		SenderPubKeyRaw: msg.SenderPubKeyRaw,
		RelayChain:      msg.RelayChain,
		PayloadSize:     uint64(len(msg.Payload)),
	})

	hash = crypto.Keccak256(bytes)

	signature, err := util.Sign(senderPrivateKey, hash)

	msg.DataWithPayloadSizeHash = hash
	msg.DataWithPayloadSizeSignature = signature

	return &msg, err
}

func (msg *Message) Relay(id wendy.NodeID, nextID wendy.NodeID, relayerPrivateKey *ecdsa.PrivateKey) error {

	currentID := msg.Sender
	var prevBlock *RelayBlock
	if len(msg.RelayChain) > 0 {
		currentID = msg.RelayChain[0].NextID
		prevBlock = &msg.RelayChain[0]
	}

	if currentID != id {
		return errors.New("The supplied ID does not match the NextID of the latest block in the Proof of Relay chain")
	}

	block, err := makeRelayBlock(id, relayerPrivateKey, nextID, prevBlock)
	if err != nil {
		return err
	}

	msg.RelayChain = append(msg.RelayChain, *block)
	return nil
}

func (msg *Message) ValidateRelayPath() error {
	if len(msg.RelayChain) == 0 {
		return errors.New("Empty relay chain")
	}

	relayLen := uint(len(msg.RelayChain))

	for i := uint(0); i < relayLen; i++ {

		relayBlock := msg.RelayChain[i]

		if i == 0 {
			// root block special validation
			err := validateRootBlockAgainstSenderData(&relayBlock, msg)
			if err != nil {
				return err
			}

		} else {

			previousBlock := msg.RelayChain[i-1]

			// normal block validation
			err := validateAgainstPrevious(&relayBlock, &previousBlock, i)
			if err != nil {
				return err
			}
		}

		// common validation for both block types
		valid, err := signatureValidation(&relayBlock)
		if err != nil {
			return err
		}

		if !valid {
			return errors.New("Relay block signature validation failed")
		}
	}

	return nil
}

func validateRootBlockAgainstSenderData(relayBlock *RelayBlock, msg *Message) error {
	if relayBlock.ID != msg.Sender {
		return errors.New("ID of the relay chain root is not the same as the sender ID")
	}

	relayKey, err := util.UnmarshalPubKey(relayBlock.PubKeyRaw)
	if err != nil {
		return err
	}

	senderKey, err := util.UnmarshalPubKey(msg.SenderPubKeyRaw)
	if err != nil {
		return err
	}

	if !util.PubKeysEqual(*relayKey, *senderKey) {
		return errors.New("Public key of the relay chain root is not the same as the sender public key")
	}

	return nil
}

func validateAgainstPrevious(relayBlock *RelayBlock, previousBlock *RelayBlock, i uint) error {
	if relayBlock.ID != previousBlock.NextID {
		return fmt.Errorf("Index %d: ID of relay block does not match NextID of previous block", i)
	}

	if relayBlock.PrevID != previousBlock.ID {
		return fmt.Errorf("Index %d: Previous ID of relay block does not match ID of previous block", i)
	}

	if !bytes.Equal(relayBlock.PrevPubKeyRaw, previousBlock.PubKeyRaw) {
		return fmt.Errorf("Index %d: Previous PubKey of relay block does not match PubKey of previous block", i)
	}

	if !bytes.Equal(relayBlock.PrevSignature, previousBlock.Signature) {
		return fmt.Errorf("Index %d: Previous signature of relay block does not match signature of previous block", i)
	}

	return nil
}

func signatureValidation(relayBlock *RelayBlock) (bool, error) {
	blockHash, err := relayBlock.BlockHash256()
	if err != nil {
		return false, err
	}

	return util.Verify(relayBlock.PubKeyRaw, blockHash[:], relayBlock.Signature), nil
}

func DefaultMessageProcessor(msg *Message) {
	//fmt.Println("Processing message with the default processor...")
	fmt.Println(string(msg.Payload))
}
