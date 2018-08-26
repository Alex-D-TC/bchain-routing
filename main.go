package main

import (
	"fmt"
	"math/big"

	"secondbit.org/wendy"

	"github.com/alex-d-tc/bchain-routing/cmd"
	"github.com/alex-d-tc/bchain-routing/eth"
	"github.com/alex-d-tc/bchain-routing/swiss"
	"github.com/alex-d-tc/bchain-routing/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//testSendRelayRequest()
	/*
		id, err := util.NodeIDFromHexForm("54e9cbfb5e6b16c5220a7468c86164b0")
		if err != nil {
			panic(err)
		}

		testGetRelayRequest(id.Base10(), big.NewInt(0))
	*/

	//testMsgRelay()
	cmd.Execute()
}

func testMsgRelay() {

	k1, err := util.LoadKeys("./RES/eth.key")
	if err != nil {
		panic(err)
	}

	k2, err := util.LoadKeys("./RES/eth2.key")
	if err != nil {
		panic(err)
	}

	sender := wendy.NodeID{0, 1}
	relayer := wendy.NodeID{6, 9}
	receiver := wendy.NodeID{1, 2}

	msg, err := swiss.MakeMessage(sender, k1, receiver, []byte("OMFGYES"))
	err = msg.Relay(sender, relayer, k1)
	if err != nil {
		panic(err)
	}

	err = msg.Relay(relayer, receiver, k2)
	if err != nil {
		panic(err)
	}

	err = msg.ValidateRelayPath()
	if err != nil {
		panic(err)
	}

	solRelay, err := swiss.MakeSolidityRelay(msg, []byte("CACAT"))
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(solRelay.Relayers); i++ {
		fmt.Println(solRelay.Relayers[i].Hex())
	}

	hash, err := swiss.IPFSStoreRelayFile(msg)
	if err != nil {
		panic(err)
	}

	fmt.Println(hash)
}

func testGetRelayRequest(id common.Address, relayId *big.Int) {
	client, err := eth.GetThreadsafeClient("https://ropsten.infura.io")
	if err != nil {
		panic(err)
	}

	contracts, err := util.ReadContractsConfig("./RES/contracts.json")
	if err != nil {
		panic(err)
	}

	relay, err := eth.GetRelayHandler(common.HexToAddress(contracts.Relay), client)
	if err != nil {
		panic(err)
	}

	res, err := relay.Relay.GetRelay(nil, id, relayId)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	/*
		raw, err := util.IPFSReadFile(string(res.IpfsRelayHash))
		if err != nil {
			panic(err)
		}

		fmt.Println(string(raw))
	*/
	pkey, err := util.UnmarshalPubKey(res.SenderPubkeyRaw)
	if err != nil {
		panic(err)
	}

	fmt.Println(crypto.PubkeyToAddress(*pkey).Hex())
	fmt.Println(string(res.IpfsRelayHash))
	fmt.Println(res.Honored)
	fmt.Println(res.SentBytes)
	fmt.Println(res.Relayers[0].Hex())
}

func testSendRelayRequest() {

	srcID := wendy.NodeID{0, 1}
	destID := wendy.NodeID{3, 4}

	contracts, err := util.ReadContractsConfig("./RES/contracts.json")
	if err != nil {
		panic(err)
	}

	client, err := eth.GetThreadsafeClient("https://ropsten.infura.io")
	if err != nil {
		panic(err)
	}

	relay, err := eth.GetRelayHandler(common.HexToAddress(contracts.Relay), client)
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

	ipfsID, err := swiss.IPFSStoreRelayFile(msg)
	if err != nil {
		panic(err)
	}

	done := make(chan uint8)

	client.SubmitTransaction(func(client *ethclient.Client) (error, bool) {
		auth, err := eth.PrepareTransactionAuth(client, k1)
		if err != nil {
			return err, false
		}

		solidityRelay, err := swiss.MakeSolidityRelay(msg, []byte(ipfsID))
		if err != nil {
			return err, false
		}

		fmt.Println("Submitting relay request")

		tx, err := relay.Relay.SubmitRelay(
			auth,
			solidityRelay.SenderEthAddress,
			solidityRelay.SenderPubKeyRaw,
			solidityRelay.Sender.Base10(),
			solidityRelay.Receiver.Base10(),
			solidityRelay.SentByteCount,
			solidityRelay.IpfsRelayHash,
			solidityRelay.Relayers)

		if err == nil {
			fmt.Println(tx.Hash().Hex())
		} else {
			fmt.Println(err)
		}

		done <- 1

		return err, false
	})

	// Wait for the transaction to be done
	<-done
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
