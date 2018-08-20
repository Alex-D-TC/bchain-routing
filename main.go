package main

import (
	"fmt"

	eth "github.com/alex-d-tc/bchain-routing/eth/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	//cmd.Execute()
}

func testEthCall() {

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

	key, _ := crypto.GenerateKey()

	fmt.Println(balanceOf)

}
