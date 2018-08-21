package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	eth "github.com/alex-d-tc/bchain-routing/eth/build-go"
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

	client, _ := makeClient("https://ropsten.infura.io")
	//printBalance(client, crypto.PubkeyToAddress(key.PublicKey))

	//testSendEth(key)
	//testEthCall(client, key)
	//testSimpleCall(client, key)
	//testContractDeploy(client, key)
	testSwissMint(client, key, 300, "0x158d29DBa40Ea09D371b23D0F20F30EedC2A4588", "0xa7ea06f6e990c26439ebd2691a47610db41580d8")
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
	auth, err := prepareTransactionAuth(client, key, 0)
	if err != nil {
		panic(err)
	}

	address, tx, instance, err := eth.DeploySimpleToken(auth, client, "Swiss", "SWS", big.NewInt(1000000000))
	if err != nil {
		panic(err)
	}

	_ = instance

	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())
}

func testSimpleCall(client *ethclient.Client, key *ecdsa.PrivateKey) {

	targetContract := common.HexToAddress("0xf480a29977435b07de7bfca98524713efcf7d7c7")
	counterInstance, err := eth.NewCounter(targetContract, client)
	if err != nil {
		panic(err)
	}

	auth, err := prepareTransactionAuth(client, key, 0)
	if err != nil {
		panic(err)
	}

	tx, err := counterInstance.Add(auth, big.NewInt(3))
	if err != nil {
		panic(err)
	}

	fmt.Println(tx.Hash().Hex())
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		panic(err)
	}

	fmt.Println(string(receipt.PostState))
}

func prepareTransactionAuth(client *ethclient.Client, key *ecdsa.PrivateKey, value int64) (*bind.TransactOpts, error) {

	/*
		nonce, err := client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(key.PublicKey))
		if err != nil {
			return nil, err
		}

		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			return nil, err
		}
	*/

	auth := bind.NewKeyedTransactor(key)

	//auth.Nonce = big.NewInt(int64(nonce))
	//auth.Value = big.NewInt(0)
	//auth.GasLimit = uint64(3000000)
	//auth.GasPrice = gasPrice

	return auth, nil
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
