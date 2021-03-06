package routingdriver

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/alex-d-tc/bchain-routing/net"
	"secondbit.org/wendy"
)

const channelBufferSize = 100

// RoutingDriver is the object used to communicate with the underlying P2P network
// It allows the sending and receiving of messages asynchronously
// TODO: Implement the bridge pattern for the RoutingDriver
type RoutingDriver struct {
	node    *wendy.Node
	cluster *wendy.Cluster

	messageBus            <-chan []byte
	killMessageProcessor  chan byte
	startMessageProcessor chan byte

	running bool

	logger *log.Logger
}

func MakeRoutingDriver(nodeID net.NodeID, localIP string, globalIP string, port int, forwardMangler func([]byte, net.NodeID) ([]byte, bool)) *RoutingDriver {
	node := wendy.NewNode(wendy.NodeID(nodeID), localIP, globalIP, "1", port)

	messageBus := make(chan []byte, channelBufferSize)
	hook := makeWendyHook(messageBus, func(msg *wendy.Message, next wendy.NodeID) bool {
		return forwardingProcessor(msg, net.NodeID(next), forwardMangler)
	})

	cluster := wendy.NewCluster(node, credentials{})
	cluster.RegisterCallback(hook)

	killProcessor := make(chan byte, 1)
	startProcessor := make(chan byte, 1)

	startProcessor <- 1

	return &RoutingDriver{
		node:                  node,
		cluster:               cluster,
		messageBus:            messageBus,
		running:               false,
		killMessageProcessor:  killProcessor,
		startMessageProcessor: startProcessor,
		logger:                log.New(os.Stdout, "Swiss routing driver ", log.Ltime|log.Ldate),
	}
}

func forwardingProcessor(msg *wendy.Message, next net.NodeID, forwardMangler func([]byte, net.NodeID) ([]byte, bool)) bool {

	payload, toSend := forwardMangler(msg.Value, next)
	msg.Value = payload

	if toSend {
		fmt.Printf("Forwarding message %s to Node %s.", msg.Key, next)
	}

	return toSend
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

func (driver *RoutingDriver) Send(destinationAddr net.NodeID, messageData []byte) error {
	message := driver.cluster.NewMessage(255, wendy.NodeID(destinationAddr), messageData)
	return driver.cluster.Send(message)
}

func (driver *RoutingDriver) SetOutput(writer io.Writer) {
	driver.logger = log.New(writer, driver.logger.Prefix(), driver.logger.Flags())
}

func (driver *RoutingDriver) processMessages(processor func([]byte)) {
	end := false

	driver.debug("Waiting until the current processor stops...")

	// Wait until the former processor is stopped
	<-driver.startMessageProcessor

	driver.debug("Driver started awaiting messages...")

	for {
		select {
		case msg := <-driver.messageBus:
			driver.debug("Driver received message")
			processor(msg)
			break
		case <-driver.killMessageProcessor:
			driver.debug("Driver received kill command")
			end = true
			break
		}

		if end {
			break
		}
	}

	// When killed, signal that another processor can begin
	driver.startMessageProcessor <- 1
}

func (driver *RoutingDriver) debug(msg string) {
	driver.logger.Println(msg)
}
