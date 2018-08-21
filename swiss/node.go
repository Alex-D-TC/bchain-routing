package swiss

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"os"

	"github.com/alex-d-tc/bchain-routing/routingdriver"
	"github.com/alex-d-tc/bchain-routing/util"
	"secondbit.org/wendy"
)

type SwissNode struct {
	ID         wendy.NodeID
	PrivateKey *ecdsa.PrivateKey

	started bool

	driver *routingdriver.RoutingDriver
	logger *log.Logger
}

func InitSwissNode(localIP string, port int, publicIP string, privKey *ecdsa.PrivateKey) *SwissNode {

	id := util.NodeIDFromStringSHA(fmt.Sprintf("%s:%d", localIP, port))

	node := &SwissNode{
		ID:         id,
		started:    false,
		logger:     log.New(os.Stdout, "Swiss node ", log.Ldate|log.Ltime),
		PrivateKey: privKey,
	}

	node.driver = routingdriver.MakeRoutingDriver(id, localIP, publicIP, port, node.forwardingProcessor)

	return node
}

func (node *SwissNode) Start(processor func(*Message)) {
	if !node.started {
		node.started = true
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
	if node.started {
		node.started = false
		node.driver.Stop()
	}
}

func (node *SwissNode) Send(destination wendy.NodeID, payload []byte) error {

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

func (node *SwissNode) SetLogger(logger *log.Logger) {
	node.logger = logger
	node.driver.SetLogger(logger)
}

func (node *SwissNode) debug(msg string) {
	node.logger.Println(msg)
}

func (node *SwissNode) forwardingProcessor(rawPayload []byte, next wendy.NodeID) ([]byte, bool) {

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
