package routingdriver

import (
	"log"
	"os"

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

	logger *log.Logger
}

func MakeRoutingDriver(nodeID wendy.NodeID, localIP string, globalIP string, port int) *RoutingDriver {
	node := wendy.NewNode(nodeID, localIP, globalIP, "1", port)

	messageBus := make(chan []byte, channelBufferSize)
	hook := &wendyHook{
		OutputChan: messageBus,
	}

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

func (driver *RoutingDriver) SetLogger(logger *log.Logger) {
	driver.logger = logger
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
