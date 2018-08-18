package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/alex-d-tc/bchain-routing/swiss"
	"github.com/alex-d-tc/bchain-routing/util"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client-start",
	Short: "Start a client node",
	Long:  "Starrrrt a cliennnnt nooooode",
	Run: func(cmd *cobra.Command, args []string) {
		runClient(clientLocalIP, clientGlobalIP, int(clientPort), clientBootstrapIP, int(clientBootstrapPort))
	},
}

var clientPort int32
var clientLocalIP string
var clientGlobalIP string
var clientBootstrapIP string
var clientBootstrapPort int32

func init() {
	flags := clientCmd.Flags()

	flags.Int32Var(&clientPort, "port", 8080, "The client port")
	flags.StringVar(&clientLocalIP, "local-ip", "127.0.0.1", "The local ip")
	flags.StringVar(&clientGlobalIP, "global-ip", "0.0.0.0", "The global ip")
	flags.StringVar(&clientBootstrapIP, "bootstrap-ip", "127.0.0.1", "The bootstrap node ip")
	flags.Int32Var(&clientBootstrapPort, "bootstrap-port", 8000, "The client bootstrap port")

	rootCmd.AddCommand(clientCmd)
}

func runClient(localIP string, publicIP string, port int, bootstrapIP string, bootstrapPort int) {

	fmt.Println("Building swiss node...")

	node := swiss.InitSwissNode(localIP, port, publicIP)

	fmt.Println("Starting routines...")

	failChan := make(chan byte)
	lineChan := make(chan string)

	// Network listener goroutine
	go func() {
		err := node.JoinAndStart(swiss.DefaultMessageProcessor, bootstrapIP, bootstrapPort)
		if err != nil {
			fmt.Println(err)
			failChan <- 1
		}
	}()

	// Stdin listener gorouting
	go func() {

		reader := bufio.NewReader(os.Stdin)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println(err)
				continue
			}
			lineChan <- line
		}

	}()

	fmt.Println(fmt.Sprintf("Client id: %s", util.NodeIDToString(node.Id)))

	for {
		fail := false
		select {
		case <-failChan:
			fail = true
		case line := <-lineChan:
			line = strings.TrimSpace(line)
			if line == "quit" {
				fail = true
				break
			}
			processCommand(line, node)
		}

		if fail {
			fmt.Println("Stopping...")
			break
		}
	}
}

func processCommand(rawLine string, node *swiss.SwissNode) {
	lineSplit := strings.Split(rawLine, " ")
	command := lineSplit[0]

	switch command {
	case "send":
		receiver := lineSplit[1]
		message := lineSplit[2]

		id, err := util.NodeIDFromHexForm(receiver)
		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(fmt.Sprintf("Sending %s to %s", message, receiver))
		node.Send(id, []byte(message))
	}

}
