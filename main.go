package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	eth "github.com/alex-d-tc/bchain-routing/eth/contracts"
	"github.com/alex-d-tc/bchain-routing/swiss"
	"github.com/alex-d-tc/bchain-routing/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//testKeys("./general.key")

	//generateEthAddress("eth.key")
	// Generated 0x4Ee601163EEc9863c6f966be984D2AF5E23185D2

	key, err := crypto.LoadECDSA("eth.key")
	if err != nil {
		panic(err)
	}
	testSendEth(key)

	//testEthCall(key)
	//cmd.Execute()
}

func readEthAddress(keyPath string) (*ecdsa.PrivateKey, error) {
	return crypto.LoadECDSA(keyPath)
}

func generateEthAddress(keyPath string) {
	key, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	fmt.Println(crypto.PubkeyToAddress(key.PublicKey).Hex())

	err = crypto.SaveECDSA(keyPath, key)
	if err != nil {
		panic(err)
	}
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

func testSendEth(key *ecdsa.PrivateKey) {
	//client, err := ethclient.Dial("https://ropsten.infura.io/v3/4bf23add04f743fab3e7bba300ea772a")
	client, err := ethclient.Dial("\\\\.\\pipe\\geth.ipc")
	if err != nil {
		panic(err)
	}

	localAddr := crypto.PubkeyToAddress(key.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), localAddr)
	if err != nil {
		panic(err)
	}

	value := big.NewInt(5000)
	gasLimit := uint64(21000)

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	destAddr := common.HexToAddress("0x158d29DBa40Ea09D371b23D0F20F30EedC2A4588")

	mulResult := gasPrice.Mul(gasPrice, big.NewInt(int64(gasLimit)))

	fmt.Println(client.BalanceAt(context.Background(), destAddr, nil))
	fmt.Println(client.BalanceAt(context.Background(), crypto.PubkeyToAddress(key.PublicKey), nil))
	fmt.Println(mulResult.Add(mulResult, value))

	tx := types.NewTransaction(nonce, destAddr, value, gasLimit, gasPrice, nil)
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, key)
	if err != nil {
		panic(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		panic(err)
	}
	fmt.Println(signedTx.Hash().Hex())
}

func testEthCall(key *ecdsa.PrivateKey) {

	// pipePath \\\\.\\pipe\\geth.ipc

	//client, err := ethclient.Dial("https://rinkeby.infura.io/")
	client, err := ethclient.Dial("\\\\.\\pipe\\geth.ipc")
	if err != nil {
		panic(err)
	}

	id, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Network id: ", id)

	swissCoinAddress := common.HexToAddress("0xd9b4a78e1aa32c4132544d1e8fdd9f4c69448b13")
	instance, err := eth.NewSwissCoin(swissCoinAddress, client)
	if err != nil {
		panic(err)
	}

	tcAddr := common.HexToAddress("0x158d29DBa40Ea09D371b23D0F20F30EedC2A4588")

	tcBalance, err := client.BalanceAt(context.Background(), tcAddr, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("Remote balance eth: ", tcBalance)

	balanceOf, err := instance.BalanceOf(nil, tcAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Destination balance SWS: ", balanceOf)

	localAddr := crypto.PubkeyToAddress(key.PublicKey)
	fmt.Println("Source address: ", localAddr.Hex())

	nonce, err := client.PendingNonceAt(context.Background(), localAddr)
	if err != nil {
		panic(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	balance, err := client.BalanceAt(context.Background(), localAddr, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Source balance eth: ", balance)

	auth := bind.NewKeyedTransactor(key)

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)
	auth.GasPrice = gasPrice

	fmt.Println("Running costs eth:  ", auth.GasPrice.Mul(auth.GasPrice, big.NewInt(int64(auth.GasLimit))))

	fmt.Println(auth.From.Hex(), localAddr.Hex())

	tx, err := instance.Transfer(auth, common.HexToAddress("0x158d29DBa40Ea09D371b23D0F20F30EedC2A4588"), big.NewInt(2000))

	if err != nil {
		panic(err)
	}

	fmt.Println(tx.Hash().Hex())

	fmt.Println(instance.BalanceOf(nil, common.HexToAddress("0x158d29DBa40Ea09D371b23D0F20F30EedC2A4588")))
	fmt.Println(instance.BalanceOf(nil, localAddr))
}
