package swiss

import "github.com/alex-d-tc/bchain-routing/net"

type IPFSRelay struct {
	Sender   net.NodeID
	Receiver net.NodeID

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
