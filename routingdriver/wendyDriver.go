package routingdriver

import (
	"fmt"

	"secondbit.org/wendy"
)

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
	//fmt.Println("Received message: ", msg)
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
