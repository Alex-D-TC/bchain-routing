package main

import (
	"crypto/sha256"
	"encoding/binary"
	"flag"
	"fmt"
	"strconv"

	"secondbit.org/wendy"
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

func nodeIDFromBytesSHA(entropySource []byte) wendy.NodeID {
	result := sha256.Sum256(entropySource)
	var id wendy.NodeID

	id[0] = binary.BigEndian.Uint64(result[:8])
	id[1] = binary.BigEndian.Uint64(result[8:16])

	return id
}

func nodeIDFromString(str string) wendy.NodeID {
	return nodeIDFromBytesSHA([]byte(str))
}

func main() {

	joinPtr, portPtr, publicIPPtr, bootstrapIPPtr, bootstrapPortPtr, _ := getFlags()

	ip := "192.168.31.113"
	id := nodeIDFromBytesSHA([]byte(ip + ":" + strconv.Itoa(*portPtr)))

	node := wendy.NewNode(id, ip, *publicIPPtr, "1", *portPtr)
	cluster := wendy.NewCluster(node, Credentials{})
	cluster.RegisterCallback(&MyApp{})

	if *joinPtr {
		fmt.Println("Joining...")
		cluster.Join(*bootstrapIPPtr, *bootstrapPortPtr)
		fmt.Println("Joined!")
	} else {
		fmt.Println("Initializing cluster...")
	}

	fmt.Println("Listening...")
	cluster.Listen()
	defer cluster.Stop()
}

type Credentials struct {
}

func (cred Credentials) Marshal() []byte {
	return make([]byte, 0)
}

func (cred Credentials) Valid(raw []byte) bool {
	return true
}
