package main

import (
	"fmt"

	"github.com/alex-d-tc/bchain-routing/routingdriver"
)

func deployRootNode(ip string, port int, publicIP string) *routingdriver.SwissNode {
	node := routingdriver.InitSwissNode(ip, port, publicIP)
	node.Start()
	return node
}

func deployNode(ip string, port int, publicIP string, boostrapIP string, bootstrapPort int) *routingdriver.SwissNode {
	node := routingdriver.InitSwissNode(ip, port, publicIP)
	node.Start()
	for {
		err := node.Join(boostrapIP, bootstrapPort)
		if err == nil {
			break
		}

		fmt.Println(err)
	}

	return node
}
