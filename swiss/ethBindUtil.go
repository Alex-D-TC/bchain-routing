package swiss

import (
	"fmt"
	"math/big"

	"github.com/alex-d-tc/bchain-routing/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"secondbit.org/wendy"
)

type SolidityRelay struct {

	// Message validation data
	SenderEthAddress common.Address
	SenderPubKeyRaw  []byte

	Sender   wendy.NodeID
	Receiver wendy.NodeID

	SentByteCount *big.Int
	Relayers      []common.Address

	// IPFS Proof of relay reference data
	IpfsRelayHash []byte
}

func MakeSolidityRelay(msg *Message, IPFSAddress []byte) (SolidityRelay, error) {
	relay := SolidityRelay{}

	senderPubkey, err := util.UnmarshalPubKey(msg.SenderPubKeyRaw)
	if err != nil {
		return SolidityRelay{}, err
	}

	relay.SenderEthAddress = crypto.PubkeyToAddress(*senderPubkey)
	relay.SenderPubKeyRaw = msg.SenderPubKeyRaw
	relay.Sender = msg.Sender
	relay.Receiver = msg.Receiver
	relay.SentByteCount = big.NewInt(int64(len(msg.Payload)))

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
