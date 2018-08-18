package swiss

import (
	"crypto/rsa"
	"fmt"
	"log"
	"os"

	"github.com/alex-d-tc/bchain-routing/routingdriver"
	"github.com/alex-d-tc/bchain-routing/util"
	"secondbit.org/wendy"
)

type SwissNode struct {
	Id         wendy.NodeID
	PrivateKey *rsa.PrivateKey

	started bool

	driver *routingdriver.RoutingDriver
	logger *log.Logger
}

func InitSwissNode(localIP string, port int, publicIP string, privKey *rsa.PrivateKey) *SwissNode {

	id := util.NodeIDFromStringSHA(fmt.Sprintf("%s:%d", localIP, port))

	return &SwissNode{
		driver:  routingdriver.MakeRoutingDriver(id, localIP, publicIP, port),
		Id:      id,
		started: false,
		logger:  log.New(os.Stdout, "Swiss node ", log.Ldate|log.Ltime),
	}
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

	message, err := MakeMessage(&node.Id, node.PrivateKey, &destination, payload)
	if err != nil {
		return err
	}

	encoder, buffer := util.MakeEncoder()
	err = encoder.Encode(*message)
	if err != nil {
		return err
	}

	return node.driver.Send(destination, buffer.Bytes())
}

func (node *SwissNode) SetLogger(logger *log.Logger) {
	node.logger = logger
	node.driver.SetLogger(logger)
}

func (node *SwissNode) debug(msg string) {
	node.logger.Println(msg)
}

func (node *SwissNode) processRaw(rawMsg []byte) (*Message, error) {
	var result Message

	decoder, buffer := util.MakeDecoder()
	buffer.Write(rawMsg)

	err := decoder.Decode(&result)
	return &result, err
}

func (node *SwissNode) processMessage(rawMsg []byte, swissProcessor func(*Message)) {
	msg, err := node.processRaw(rawMsg)
	if err != nil {
		fmt.Println(err)
	} else {
		swissProcessor(msg)
	}
}
