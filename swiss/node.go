package swiss

import (
	"crypto/ecdsa"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/alex-d-tc/bchain-routing/eth"
	"github.com/alex-d-tc/bchain-routing/net"
	"github.com/alex-d-tc/bchain-routing/routingdriver"
	"github.com/alex-d-tc/bchain-routing/util"
)

type SwissNode struct {
	ID         net.NodeID
	PrivateKey *ecdsa.PrivateKey

	Started bool

	driver *routingdriver.RoutingDriver
	logger *log.Logger

	client *eth.ThreadsafeClient

	coin  eth.CoinContract
	relay eth.RelayContract
}

func InitSwissNode(localIP string, port int, publicIP string, privKey *ecdsa.PrivateKey, client *eth.ThreadsafeClient, relay eth.RelayContract, coin eth.CoinContract) *SwissNode {

	id := util.NodeIDFromStringSHA(fmt.Sprintf("%s:%d", localIP, port))

	node := &SwissNode{
		ID:         id,
		Started:    false,
		logger:     log.New(os.Stdout, "Swiss node ", log.Ldate|log.Ltime),
		PrivateKey: privKey,
		client:     client,
		relay:      relay,
		coin:       coin,
	}

	node.driver = routingdriver.MakeRoutingDriver(id, localIP, publicIP, port, node.forwardingProcessor)

	return node
}

func (node *SwissNode) Start(processor func(*Message)) {
	if !node.Started {
		node.Started = true
		node.driver.Start(func(rawBytes []byte) {
			node.processMessage(rawBytes, processor)
		})
	}
}

func (node *SwissNode) JoinAndStart(processor func(*Message), bootstrapIP string, bootstrapPort int) error {
	err := node.driver.Join(bootstrapIP, bootstrapPort)
	if err == nil {
		node.Start(processor)
	}
	return err
}

func (node *SwissNode) Terminate() {
	if node.Started {
		node.Started = false
		node.driver.Stop()
	}
}

func (node *SwissNode) Send(destination net.NodeID, payload []byte) error {

	message, err := MakeMessage(node.ID, node.PrivateKey, destination, payload)
	if err != nil {
		return err
	}

	encodingResult, err := util.GobEncode(*message)
	if err != nil {
		return err
	}

	return node.driver.Send(destination, encodingResult)
}

func (node *SwissNode) SetOutput(writer io.Writer) {
	node.logger = log.New(writer, node.logger.Prefix(), node.logger.Flags())
	node.driver.SetOutput(writer)
}

func (node *SwissNode) debug(msg ...interface{}) {
	node.logger.Println(msg...)
}

func (node *SwissNode) forwardingProcessor(rawPayload []byte, next net.NodeID) ([]byte, bool) {

	// Message decoding from raw data

	var msg Message
	err := util.GobDecode(rawPayload, &msg)
	if err != nil {
		node.debug(fmt.Sprintf("%s", err))
		return rawPayload, false
	}

	// Relaying and Validation

	err = msg.Relay(node.ID, next, node.PrivateKey)
	if err != nil {
		node.debug(fmt.Sprintf("%s", err))
		return rawPayload, false
	}

	err = msg.ValidateRelayPath()
	if err != nil {
		node.debug(fmt.Sprintf("%s", err))
		return rawPayload, false
	}

	// Last node on the route. Send relay payment request to the blockchain
	if next == msg.Receiver {
		node.client.SubmitTransaction(func(client *ethclient.Client) (error, bool) {

			auth, err := eth.PrepareTransactionAuth(client, node.PrivateKey)
			if err != nil {
				node.debug(err)
				return err, false
			}

			relay, err := MakeIPFSRelay(&msg)
			if err != nil {
				return err, false
			}

			// Store data in ipfs beforehand
			ipfsID, err := IPFSStoreRelayFile(relay)
			if err != nil {
				node.debug(err)
				return err, false
			}

			solidityRelay, err := MakeSolidityRelay(&msg, []byte(ipfsID))
			if err != nil {
				node.debug(err)
				return err, false
			}

			node.debug("Submitting relay request")

			tran, err := node.relay.Relay.SubmitRelay(
				auth,
				solidityRelay.SenderEthAddress,
				solidityRelay.SenderPubKeyRaw,
				solidityRelay.Sender.Base10(),
				solidityRelay.Receiver.Base10(),
				solidityRelay.SentByteCount,
				solidityRelay.IpfsRelayHash,
				solidityRelay.Relayers)

			if err != nil {
				node.debug(err)
			} else {
				node.debug("Relay request transaction: ", tran.Hash().Hex())
			}

			return err, false
		})
	}

	// Message encoding to raw data

	encoded, err := util.GobEncode(msg)
	if err != nil {
		node.debug(fmt.Sprintf("%s", err))
		return rawPayload, false
	}

	return encoded, true
}

func (node *SwissNode) processRaw(rawMsg []byte) (*Message, error) {
	var result Message
	err := util.GobDecode(rawMsg, &result)
	return &result, err
}

func (node *SwissNode) processMessage(rawMsg []byte, swissProcessor func(*Message)) {
	msg, err := node.processRaw(rawMsg)
	if err != nil {
		node.logger.Println(err)
	} else {
		swissProcessor(msg)
	}
}
