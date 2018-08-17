package routingdriver

import (
	"fmt"

	"secondbit.org/wendy"
)

const channelBufferSize = 100

type RoutingDriver struct {
	node    *wendy.Node
	cluster *wendy.Cluster

	messageBus            <-chan []byte
	killMessageProcessor  chan byte
	startMessageProcessor chan byte

	running bool
}

func MakeRoutingDriver(nodeID wendy.NodeID, localIP string, globalIP string, port int) *RoutingDriver {
	node := wendy.NewNode(nodeID, localIP, globalIP, "1", port)

	channel := make(chan []byte, channelBufferSize)
	hook := &wendyHook{channel}

	cluster := wendy.NewCluster(node, credentials{})
	cluster.RegisterCallback(hook)

	killProcessor := make(chan byte)
	startProcessor := make(chan byte)

	killProcessor <- 1
	startProcessor <- 1

	return &RoutingDriver{
		node:                  node,
		cluster:               cluster,
		messageBus:            channel,
		running:               false,
		killMessageProcessor:  killProcessor,
		startMessageProcessor: startProcessor,
	}
}

func (driver *RoutingDriver) processMessages(processor func([]byte)) {
	end := false

	// Wait until the former processor is stopped
	<-driver.startMessageProcessor

	for {
		select {
		case msg := <-driver.messageBus:
			processor(msg)
		case <-driver.killMessageProcessor:
			end = true
		}

		if end {
			break
		}
	}

	// When killed, signal that another processor can begin
	driver.startMessageProcessor <- 1
}

func (driver *RoutingDriver) Join(bootstrapIP string, bootstrapPort int) error {
	return driver.cluster.Join(bootstrapIP, bootstrapPort)
}

func (driver *RoutingDriver) Start(processor func([]byte)) {
	if !driver.running {
		driver.running = true

		go driver.processMessages(processor)
		driver.cluster.Listen()
	}
}

func (driver *RoutingDriver) Stop() {
	if driver.running {
		driver.running = false

		// Stop the currently running processor
		driver.killMessageProcessor <- 1

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
	fmt.Println(node.ID[0], node.ID[1])
}

type credentials struct {
}

func (cred credentials) Marshal() []byte {
	return make([]byte, 0)
}

func (cred credentials) Valid(raw []byte) bool {
	return true
}
