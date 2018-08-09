package main

import (
	"flag"

	"github.com/alex-d-tc/bchain-routing/routingdriver"
)

func getFlags() (joinPtr *bool, portPtr *int, publicIPPtr *string, bootstrapIPPtr *string, bootstrapPortPtr *int, logFilePtr *string) {
	joinPtr = flag.Bool("join", false, "join the cluster or make a new one?")
	portPtr = flag.Int("port", 8080, "The port to listen on")
	publicIPPtr = flag.String("public-ip", "", "The public ip")
	bootstrapIPPtr = flag.String("bootstrap-ip", "", "The destination ip if joining a network")
	bootstrapPortPtr = flag.Int("bootstrap-port", -1, "The destination port if joining the network")
	logFilePtr = flag.String("log-file", "", "The Log file path. Leave empty for stdout")

	flag.Parse()
	return
}

func main() {

	joinPtr, portPtr, publicIPPtr, bootstrapIPPtr, bootstrapPortPtr, _ := getFlags()

	ip := ""

	node := routingdriver.InitSwissNode(ip, *portPtr, *publicIPPtr)

	if *joinPtr {
		node.Join(*bootstrapIPPtr, *bootstrapPortPtr)
	}

	node.Start()
}
