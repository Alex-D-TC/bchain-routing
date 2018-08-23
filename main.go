package main

import (
	"fmt"

	"secondbit.org/wendy"

	"github.com/alex-d-tc/bchain-routing/eth"
	"github.com/alex-d-tc/bchain-routing/eth/build-go"
	"github.com/alex-d-tc/bchain-routing/swiss"
	"github.com/alex-d-tc/bchain-routing/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	testSendRelayRequest()
	//cmd.Execute()
}

func testSendRelayRequest() {

	srcID := wendy.NodeID{0, 1}
	destID := wendy.NodeID{3, 4}

	contracts, err := util.ReadContractsConfig("./RES/contracts.json")
	if err != nil {
		panic(err)
	}

	client, err := eth.GetClient("https://ropsten.infura.io")
	if err != nil {
		panic(err)
	}

	relay, err := ethBind.NewRelayHandler(common.HexToAddress(contracts.Relay), client)
	if err != nil {
		panic(err)
	}

	k1, err := util.LoadKeys("./RES/eth.key")
	if err != nil {
		panic(err)
	}

	msg, err := swiss.MakeMessage(srcID, k1, destID, []byte{1, 2, 3})
	if err != nil {
		panic(err)
	}

	err = msg.Relay(srcID, destID, k1)
	if err != nil {
		panic(err)
	}

	fmt.Println(crypto.PubkeyToAddress(k1.PublicKey).Hex())

	fmt.Println("Submitting relay request")

	auth, err := eth.PrepareTransactionAuth(client, k1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Preparing data...")

	solRelay, err := swiss.MakeSolidityRelay(msg)
	if err != nil {
		fmt.Println(err)
	}

	solRelay.DebugPrint()

	tran, err := relay.SubmitRelay(auth,
		solRelay.SentBytes,
		solRelay.SentBytesHash,
		solRelay.SentBytesSignature,
		solRelay.SenderPublicKey,
		solRelay.IDS,
		solRelay.Keys,
		solRelay.Signatures,
		solRelay.PorRawHash)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(tran.Hash().Hex())
	}
}

func testMsg() {

	sourceID := wendy.NodeID{1, 2}
	destID := wendy.NodeID{0, 1}

	key, err := crypto.LoadECDSA("eth.key")
	if err != nil {
		panic(err)
	}

	msg, err := swiss.MakeMessage(sourceID, key, destID, []byte{1, 3, 4})
	if err != nil {
		panic(err)
	}

	encoded, err := util.GobEncode(msg)
	if err != nil {
		panic(err)
	}

	//fmt.Println(msg)

	var resMsg swiss.Message
	err = util.GobDecode(encoded, &resMsg)
	if err != nil {
		panic(err)
	}

	//fmt.Println(resMsg)

	err = msg.Relay(sourceID, destID, key)
	if err != nil {
		panic(err)
	}

	err = msg.ValidateRelayPath()
	if err != nil {
		panic(err)
	}
}
