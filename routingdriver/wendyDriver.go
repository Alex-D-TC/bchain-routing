package routingdriver

import (
	"log"
	"os"

	"secondbit.org/wendy"
)

/// WENDY DRIVER ///
type wendyHook struct {
	OutputChan chan<- []byte
	onForward  func(*wendy.Message, wendy.NodeID) bool

	logger *log.Logger
}

func makeWendyHook(outputChan chan<- []byte, onForward func(*wendy.Message, wendy.NodeID) bool) *wendyHook {
	return &wendyHook{
		OutputChan: outputChan,
		onForward:  onForward,
		logger:     log.New(os.Stdout, "Wendy Hook: ", log.Ltime|log.Ldate),
	}
}

func (app *wendyHook) OnDeliver(msg wendy.Message) {
	//app.debug("Received message: ", msg)
	app.OutputChan <- msg.Value
}

func (app *wendyHook) OnForward(msg *wendy.Message, next wendy.NodeID) bool {
	return app.onForward(msg, next)
}

func (app *wendyHook) OnError(err error) {
	panic(err.Error())
}

func (app *wendyHook) OnNewLeaves(leaves []*wendy.Node) {
	//app.debug("Leaf set changed: ", leaves)
}

func (app *wendyHook) OnNodeJoin(node wendy.Node) {
	app.debug("Node joined: ", node.ID)
}

func (app *wendyHook) OnNodeExit(node wendy.Node) {
	app.debug("Node left: ", node.ID)
}

func (app *wendyHook) OnHeartbeat(node wendy.Node) {
	app.debug("Received heartbeat from ", node.ID)
	app.debug(node.ID[0], node.ID[1])
}

func (app *wendyHook) debug(msg ...interface{}) {
	app.logger.Println(msg...)
}

type credentials struct {
}

func (cred credentials) Marshal() []byte {
	return make([]byte, 0)
}

func (cred credentials) Valid(raw []byte) bool {
	return true
}
