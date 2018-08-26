package swiss

import (
	"secondbit.org/wendy"
)

type IPFSRelay struct {
	Sender   wendy.NodeID
	Receiver wendy.NodeID

	SenderPubKeyRaw []byte

	RelayChain []RelayBlock

	PayloadByteCount uint64

	DataHash      []byte
	DataSignature []byte
}

func MakeIPFSRelay(msg *Message) (IPFSRelay, error) {

	relay := IPFSRelay{
		Sender:           msg.Sender,
		Receiver:         msg.Receiver,
		SenderPubKeyRaw:  msg.SenderPubKeyRaw,
		RelayChain:       msg.RelayChain,
		PayloadByteCount: uint64(len(msg.Payload)),
		DataHash:         msg.DataWithPayloadSizeHash,
		DataSignature:    msg.DataWithPayloadSizeSignature,
	}

	return relay, nil
}
