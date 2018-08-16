package routingdriver

import (
	"fmt"

	"secondbit.org/wendy"
)

const channelBufferSize = 100

type RoutingDriver struct {
	node       *wendy.Node
	cluster    *wendy.Cluster
	messageBus <-chan []byte
	running    bool
}

func MakeRoutingDriver(nodeID wendy.NodeID, localIP string, globalIP string, port int) *RoutingDriver {
	node := wendy.NewNode(nodeID, localIP, globalIP, "1", port)

	channel := make(chan []byte, channelBufferSize)
	hook := &wendyHook{channel}

	cluster := wendy.NewCluster(node, credentials{})
	cluster.RegisterCallback(hook)

	return &RoutingDriver{
		node:       node,
		cluster:    cluster,
		messageBus: channel,
		running:    false,
	}
}

func (driver *RoutingDriver) ProcessMessages(processor func([]byte)) {
	for {
		processor(<-driver.messageBus)
	}
}

func (driver *RoutingDriver) Join(bootstrapIP string, bootstrapPort int) error {
	return driver.cluster.Join(bootstrapIP, bootstrapPort)
}

func (driver *RoutingDriver) Start() {
	if !driver.running {
		driver.running = true
		driver.cluster.Listen()
	}
}

func (driver *RoutingDriver) Stop() {
	if driver.running {
		driver.running = false
		driver.cluster.Stop()
	}
}

func (driver *RoutingDriver) Send(destinationAddr wendy.NodeID, messageData []byte) error {
	message := driver.cluster.NewMessage(255, destinationAddr, messageData)
	return driver.cluster.Send(message)
}

/// WENDY DRIVER ///
type wendyHook struct {
	OutputChan chan<- []byte
}

func makeWendyHook(outputChan chan<- []byte) *wendyHook {
	return &wendyHook{
		OutputChan: outputChan,
	}
}

func (app *wendyHook) OnDeliver(msg wendy.Message) {
	fmt.Println("Received message: ", msg)
	app.OutputChan <- msg.Value
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

type credentials struct {
}

func (cred credentials) Marshal() []byte {
	return make([]byte, 0)
}

func (cred credentials) Valid(raw []byte) bool {
	return true
}
