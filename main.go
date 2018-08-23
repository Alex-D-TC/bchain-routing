package main

import (
	"fmt"

	"secondbit.org/wendy"

	"github.com/alex-d-tc/bchain-routing/swiss"
	"github.com/alex-d-tc/bchain-routing/util"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//testKeys("./general.key")

	//generateEthAddress("eth.key")
	// Generated 0x4Ee601163EEc9863c6f966be984D2AF5E23185D2
	//client, _ := makeClient("https://ropsten.infura.io")
	//printBalance(client, crypto.PubkeyToAddress(key.PublicKey))

	//testSendEth(key)
	//testEthCall(client, key)
	//testSimpleCall(client, key)
	//testContractDeploy(client, key)
	//testSwissMint(client, key, 300, "0x158d29DBa40Ea09D371b23D0F20F30EedC2A4588", "0xa7ea06f6e990c26439ebd2691a47610db41580d8")
	//testMsg()
	//cmd.Execute()
	key, err := util.LoadKeys("./eth.key")
	if err != nil {
		panic(err)
	}

	fmt.Println("Using address: ", crypto.PubkeyToAddress(key.PublicKey).Hex())

	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting test...")
	testEvents("0x9E313A60637003B0829A7A9CA6133b2cF95E9C5c", client, key)
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
