package util

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	eth "github.com/alex-d-tc/bchain-routing/eth/build-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

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

func makeClient(rawURL string) (*ethclient.Client, error) {
	return ethclient.Dial(rawURL)
}

func printBalance(client *ethclient.Client, addr common.Address) {
	fmt.Println(client.BalanceAt(context.Background(), addr, nil))
}

func testSendEth(key *ecdsa.PrivateKey) {
	client, err := ethclient.Dial("https://ropsten.infura.io/")
	//client, err := ethclient.Dial("\\\\.\\pipe\\geth.ipc")
	if err != nil {
		panic(err)
	}

	localAddr := crypto.PubkeyToAddress(key.PublicKey)

	nonce, err := client.PendingNonceAt(context.Background(), localAddr)
	if err != nil {
		panic(err)
	}

	value := big.NewInt(2000)
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

func testContractDeploy(client *ethclient.Client, key *ecdsa.PrivateKey) {

	address, tx, instance, err := eth.DeploySimpleToken(bind.NewKeyedTransactor(key), client, "Swiss", "SWS", big.NewInt(1000000000))
	if err != nil {
		panic(err)
	}

	_ = instance

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())
}

func testSwissMint(client *ethclient.Client, key *ecdsa.PrivateKey, value uint, destination string, swissContract string) {

	// pipePath \\\\.\\pipe\\geth.ipc

	id, err := client.NetworkID(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Network id: ", id)

	swissCoinAddress := common.HexToAddress(swissContract)
	instance, err := eth.NewSimpleToken(swissCoinAddress, client)
	if err != nil {
		panic(err)
	}

	tcAddr := common.HexToAddress(destination)

	sourceBalanceOf, err := instance.BalanceOf(nil, crypto.PubkeyToAddress(key.PublicKey))
	if err != nil {
		panic(err)
	}
	fmt.Println("Source balance SWS: ", sourceBalanceOf)

	balanceOf, err := instance.BalanceOf(nil, tcAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Destination balance SWS: ", balanceOf)

	localAddr := crypto.PubkeyToAddress(key.PublicKey)
	fmt.Println("Source address: ", localAddr.Hex())

	tcBalance, err := client.BalanceAt(context.Background(), tcAddr, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Remote balance wei: ", tcBalance)

	balance, err := client.BalanceAt(context.Background(), localAddr, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Source balance wei: ", balance)

	/*
		auth, err := prepareTransactionAuth(client, key, 0)
		if err != nil {
			panic(err)
		}
	*/

	//fmt.Println("Running costs wei:  ", auth.GasPrice.Mul(auth.GasPrice, big.NewInt(int64(auth.GasLimit))))

	//fmt.Println(auth.From.Hex(), localAddr.Hex())

	val := big.NewInt(int64(value))
	tx, err := instance.Mint(bind.NewKeyedTransactor(key), tcAddr, val)
	if err != nil {
		panic(err)
	}

	fmt.Println(tx.Hash().Hex())

	fmt.Println(instance.BalanceOf(nil, tcAddr))
	fmt.Println(instance.BalanceOf(nil, localAddr))
}
