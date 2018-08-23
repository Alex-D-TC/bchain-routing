package swiss

import (
	"fmt"
	"math/big"

	"github.com/alex-d-tc/bchain-routing/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type SolidityRelay struct {

	// Message validation data
	Sender             *big.Int
	SentBytes          *big.Int
	SentBytesHash      []byte
	SentBytesSignature []byte
	SenderPubKey       []byte

	// IPFS Proof of relay reference data
	IpfsRelayHash []byte
	Relayers      []common.Address
}

func MakeSolidityRelay(msg *Message, IPFSAddress []byte) (SolidityRelay, error) {
	relay := SolidityRelay{}

	relay.Sender = msg.Sender.Base10()
	relay.SentBytes = big.NewInt(int64(len(msg.Payload)))
	relay.SentBytesHash = msg.ByteCountHash
	relay.SentBytesSignature = msg.ByteCountSignature
	relay.SenderPubKey = msg.SenderPubKeyRaw
	relay.IpfsRelayHash = IPFSAddress

	for i := 0; i < len(msg.RelayChain); i++ {
		relayBlock := msg.RelayChain[i]

		pubkey, err := util.UnmarshalPubKey(relayBlock.PubKeyRaw)
		if err != nil {
			return SolidityRelay{}, err
		}

		relay.Relayers = append(relay.Relayers, crypto.PubkeyToAddress(*pubkey))
	}

	return relay, nil
}

func (block SolidityRelay) DebugPrint() {
	json, err := util.JSONEncode(block)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))
}
