package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/alex-d-tc/bchain-routing/eth"
	"github.com/alex-d-tc/bchain-routing/swiss"
	"github.com/alex-d-tc/bchain-routing/util"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client-start",
	Short: "Start a client node",
	Long:  "Starrrrt a cliennnnt nooooode",
	Run: func(cmd *cobra.Command, args []string) {
		runClient(
			clientLocalIP,
			clientGlobalIP,
			int(clientPort),
			clientBootstrapIP,
			int(clientBootstrapPort),
			clientKeyPath,
			clientClientPath,
			clientConfigPath)
	},
}

var clientPort int32
var clientLocalIP string
var clientGlobalIP string
var clientBootstrapIP string
var clientBootstrapPort int32
var clientKeyPath string
var clientClientPath string
var clientConfigPath string

func init() {
	flags := clientCmd.Flags()

	flags.Int32Var(&clientPort, "port", 8080, "The client port")
	flags.StringVar(&clientLocalIP, "local-ip", "127.0.0.1", "The local ip")
	flags.StringVar(&clientGlobalIP, "global-ip", "0.0.0.0", "The global ip")
	flags.StringVar(&clientBootstrapIP, "bootstrap-ip", "127.0.0.1", "The bootstrap node ip")
	flags.Int32Var(&clientBootstrapPort, "bootstrap-port", 8000, "The client bootstrap port")
	flags.StringVar(&clientKeyPath, "key", "", "The path to the private key file")
	flags.StringVar(&clientClientPath, "conn", "https://ropsten.infura.io/", "The url to which the ethereum client connects to the network")
	flags.StringVar(&clientConfigPath, "config", "", "The config path")

	rootCmd.AddCommand(clientCmd)
}

func runClient(localIP string, publicIP string, port int, bootstrapIP string, bootstrapPort int, keyPath string, clientURL string, configPath string) {

	config, err := util.ReadContractsConfig(configPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("Building swiss node...")

	privKey, err := util.LoadKeys(keyPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Connecting to ", clientURL)

	client, err := eth.GetThreadsafeClient(clientURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	relay, err := eth.GetRelayHandler(common.HexToAddress(config.Relay), client)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	coin, err := eth.GetSwissCoin(common.HexToAddress(config.Swiss), client)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	node := swiss.InitSwissNode(localIP, port, publicIP, privKey, client, relay, coin)
	swissClient := swiss.MakeClient(node)

	fmt.Println("Starting routines...")

	failChan := make(chan byte)
	lineChan := make(chan string)

	// Network listener goroutine
	go func() {
		err := swissClient.JoinAndStart(swiss.DefaultMessageProcessor, bootstrapIP, bootstrapPort)
		if err != nil {
			fmt.Println(err)
			failChan <- 1
		}
	}()

	// Stdin listener goroutine
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

	fmt.Println(fmt.Sprintf("Client id: %s", util.NodeIDToString(node.ID)))

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
			swissClient.Terminate()
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
		err = node.Send(id, []byte(message))
		if err != nil {
			fmt.Println(err)
		}
	}

}
