package routingdriver

import (
	"fmt"

	"secondbit.org/wendy"
)

type wendyHook struct {
}

func (app *wendyHook) OnDeliver(msg wendy.Message) {
	fmt.Println("Received message: ", msg)
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
