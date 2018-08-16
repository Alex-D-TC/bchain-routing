package swiss

import (
	"fmt"
	"strconv"

	"github.com/alex-d-tc/bchain-routing/routingdriver"
	"github.com/alex-d-tc/bchain-routing/util"
	"secondbit.org/wendy"
)

type SwissNode struct {
	driver *routingdriver.RoutingDriver
	id     wendy.NodeID
}

func InitSwissNode(localIP string, port int, publicIP string) *SwissNode {

	id := util.NodeIDFromStringSHA(fmt.Sprintf("%s:%d", localIP, strconv.Itoa(port)))

	return &SwissNode{
		driver: routingdriver.MakeRoutingDriver(id, localIP, publicIP, port),
		id:     id,
	}
}

func (node *SwissNode) Start() {
	node.driver.Start()
}

func (node *SwissNode) JoinAndStart(bootstrapIP string, bootstrapPort int) error {
	err := node.driver.Join(bootstrapIP, bootstrapPort)
	if err == nil {
		node.driver.Start()
	}
	return err
}

func (node *SwissNode) Terminate() {
	node.driver.Stop()
}

func (node *SwissNode) Send(destination wendy.NodeID, message *Message) error {
	return node.driver.Send(destination, message.ToBytes())
}

func (node *SwissNode) processRaw(rawMsg []byte) (*Message, error) {
	return MessageFromBytes(rawMsg)
}

func (node *SwissNode) processMessage(rawMsg []byte) {
	msg, err := node.processRaw(rawMsg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(msg.rawString)
	}
}
