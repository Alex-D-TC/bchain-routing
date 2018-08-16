package routingdriver

import (
	"fmt"
	"strconv"

	"github.com/alex-d-tc/bchain-routing/util"
	"secondbit.org/wendy"
)

type Credentials struct {
}

func (cred Credentials) Marshal() []byte {
	return make([]byte, 0)
}

func (cred Credentials) Valid(raw []byte) bool {
	return true
}

type SwissNode struct {
	cluster *wendy.Cluster
	started bool
}

func InitSwissNode(localIP string, port int, publicIP string) *SwissNode {
	id := util.NodeIDFromStringSHA(localIP + ":" + strconv.Itoa(port))
	node := wendy.NewNode(id, localIP, publicIP, "1", port)

	cluster := wendy.NewCluster(node, Credentials{})
	cluster.RegisterCallback(&wendyHook{})

	return &SwissNode{
		cluster: cluster,
	}
}

func (node *SwissNode) Start() {
	if !node.started {
		node.started = true
		go node.cluster.Listen()
	}
}

func (node *SwissNode) Terminate() {
	if node.started {
		node.cluster.Stop()
	}
}

type SwissMsg struct {
}

func SwissMsgFromBytes([]byte) (*SwissMsg, error) {
	return nil, nil
}

func (msg *SwissMsg) ToBytes() []byte {
	return []byte{}
}

type wendyHook struct {
	outputChan chan<- SwissMsg
}

func makeWendyHook(outputChan chan<- SwissMsg) *wendyHook {
	return &wendyHook{
		outputChan: outputChan,
	}
}

func (app *wendyHook) OnDeliver(msg wendy.Message) {
	fmt.Println("Received message: ", msg)
	swissMsg, err := SwissMsgFromBytes(msg.Value)
	if err != nil {
		app.outputChan <- *swissMsg
	}
}

func (app *wendyHook) OnForward(msg *wendy.Message, next wendy.NodeID) bool {
	fmt.Printf("Forwarding message %s to Node %s.", msg.Key, next)
	return true // return false if you don't want the message forwarded
}

func (app *wendyHook) OnError(err error) {
	panic(err.Error())
}

func (app *wendyHook) OnNewLeaves(leaves []*wendy.Node) {
	fmt.Println("Leaf set changed: ", leaves)
}

func (app *wendyHook) OnNodeJoin(node wendy.Node) {
	fmt.Println("Node joined: ", node.ID)
}

func (app *wendyHook) OnNodeExit(node wendy.Node) {
	fmt.Println("Node left: ", node.ID)
}

func (app *wendyHook) OnHeartbeat(node wendy.Node) {
	fmt.Println("Received heartbeat from ", node.ID)
}
