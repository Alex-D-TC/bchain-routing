package swiss

import (
	"fmt"

	"github.com/alex-d-tc/bchain-routing/routingdriver"
	"github.com/alex-d-tc/bchain-routing/util"
	"secondbit.org/wendy"
)

type SwissNode struct {
	driver  *routingdriver.RoutingDriver
	Id      wendy.NodeID
	started bool
}

func InitSwissNode(localIP string, port int, publicIP string) *SwissNode {

	id := util.NodeIDFromStringSHA(fmt.Sprintf("%s:%d", localIP, port))

	return &SwissNode{
		driver:  routingdriver.MakeRoutingDriver(id, localIP, publicIP, port),
		Id:      id,
		started: false,
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

func (node *SwissNode) Send(destination wendy.NodeID, message *Message) error {
	return node.driver.Send(destination, message.ToBytes())
}

func (node *SwissNode) processRaw(rawMsg []byte) (*Message, error) {
	return MessageFromBytes(rawMsg)
}

func (node *SwissNode) processMessage(rawMsg []byte, swissProcessor func(*Message)) {
	msg, err := node.processRaw(rawMsg)
	if err != nil {
		fmt.Println(err)
	} else {
		swissProcessor(msg)
	}
}
