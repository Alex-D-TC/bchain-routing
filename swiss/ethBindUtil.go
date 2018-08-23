package swiss

import "math/big"

type SolidityRelay struct {

	// Message validation data
	SentBytes          *big.Int
	SentBytesHash      []byte
	SentBytesSignature []byte
	SenderPublicKey    []byte

	// Proof of relay 'compressed' data
	IDS        [3][]*big.Int
	Keys       [2][][]byte
	Signatures [2][][]byte
	PorRawHash [][]byte
}

func MakeSolidityRelay(msg *Message) (SolidityRelay, error) {
	relay := SolidityRelay{}

	relay.SentBytes = big.NewInt(int64(len(msg.Payload)))
	relay.SentBytesHash = msg.ByteCountHash
	relay.SentBytesSignature = msg.ByteCountSignature
	relay.SenderPublicKey = msg.SenderPubKeyRaw

	IDS := []*big.Int{}
	prevIDS := []*big.Int{}
	nextIDS := []*big.Int{}

	pubkey := [][]byte{}
	prevPubkey := [][]byte{}

	porSignature := [][]byte{}
	porPrevSignature := [][]byte{}

	porRawHash := [][]byte{}

	for i := 0; i < len(msg.RelayChain); i++ {

		relayBlock := msg.RelayChain[i]

		IDS = append(IDS, relayBlock.ID.Base10())
		prevIDS = append(prevIDS, relayBlock.PrevID.Base10())
		nextIDS = append(nextIDS, relayBlock.NextID.Base10())

		pubkey = append(pubkey, relayBlock.PubKeyRaw)
		prevPubkey = append(prevPubkey, relayBlock.PrevPubKeyRaw)

		porSignature = append(porSignature, relayBlock.Signature)
		porPrevSignature = append(porPrevSignature, relayBlock.PrevSignature)

		rawHash, err := relayBlock.BlockHash256()
		if err != nil {
			return SolidityRelay{}, err
		}

		porRawHash = append(porRawHash)
	}

	relay.IDS = [3][]*big.Int{IDS, prevIDS, nextIDS}
	relay.Keys = [2][][]byte{pubkey, prevPubkey}
	relay.Signatures = [2][][]byte{porSignature, porPrevSignature}
	relay.PorRawHash = porRawHash

	return relay, nil
}
