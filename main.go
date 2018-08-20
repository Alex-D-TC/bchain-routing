package main

import (
	"crypto/ecdsa"
	"fmt"

	eth "github.com/alex-d-tc/bchain-routing/eth/contracts"
	"github.com/alex-d-tc/bchain-routing/swiss"
	"github.com/alex-d-tc/bchain-routing/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//testKeys("./general.key")
	key, err := util.LoadKeys("./general.key")
	if err != nil {
		panic(err)
	}
	testEthCall(key)
	//cmd.Execute()
}

func testKeys(keyPath string) {
	k1, err := util.GenerateECDSAKey()
	if err != nil {
		fmt.Println(err)
	}

	util.WriteKeys(keyPath, k1)
	k2, err := util.LoadKeys(keyPath)
	if err != nil {
		fmt.Println(err)
	}

	msg, err := swiss.MakeMessage([2]uint64{0, 1}, k2, [2]uint64{1, 2}, []byte{1, 2, 3})
	if err != nil {
		fmt.Println(err)
	}

	json, _ := util.JSONEncode(msg)

	var message swiss.Message
	fmt.Println(util.JSONDecode(json, &message))

	fmt.Println(util.PubKeysEqual(k1.PublicKey, k2.PublicKey))
}

func testEthCall(key ecdsa.PrivateKey) {

	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		panic(err)
	}

	address := common.HexToAddress("0xe321cf3d74f46f26966babd585aa6e881ebbf652")
	instance, err := eth.NewSwissCoin(address, client)
	if err != nil {
		panic(err)
	}

	balanceOf, err := instance.BalanceOf(nil, common.HexToAddress("0x158d29DBa40Ea09D371b23D0F20F30EedC2A4588"))
	if err != nil {
		panic(err)
	}

	fmt.Println(crypto.PublicKeyToAddress(key.PublicKey))

	fmt.Println(balanceOf)
}
