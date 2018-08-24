package cmd

import (
	"fmt"
	"os"

	"github.com/alex-d-tc/bchain-routing/eth"
	"github.com/alex-d-tc/bchain-routing/swiss"
	"github.com/alex-d-tc/bchain-routing/util"
	"github.com/ethereum/go-ethereum/common"

	"github.com/spf13/cobra"
)

var rootLocalIP string
var rootGlobalIP string
var rootPort int32
var rootKeyPath string
var rootClientPath string
var rootConfigPath string

var clusterStart = &cobra.Command{
	Use:   "cluster-start",
	Short: "Swiss cluster deploy",
	Long:  "Swisssssss clusterrrr deployyyy",
	Run: func(cmd *cobra.Command, args []string) {
		runCluster(
			rootLocalIP,
			rootGlobalIP,
			int(rootPort),
			rootKeyPath,
			rootClientPath,
			rootConfigPath)
	},
}

func init() {
	flags := clusterStart.Flags()

	flags.StringVar(&rootLocalIP, "local-ip", "127.0.0.1", "The local ip address")
	flags.StringVar(&rootGlobalIP, "global-ip", "0.0.0.0", "The global ip address")
	flags.Int32Var(&rootPort, "port", 8000, "The cluster port")
	flags.StringVar(&rootKeyPath, "key", "", "The path to the private key file")
	flags.StringVar(&rootClientPath, "conn", "https://ropsten.infura.io/", "The url to which the ethereum client connects to the network")
	flags.StringVar(&rootConfigPath, "config", "", "The config path")

	rootCmd.AddCommand(clusterStart)
}

func runCluster(localIP string, globalIP string, port int, keyPath string, clientURL string, configPath string) {

	contracts, err := util.ReadContractsConfig(configPath)
	if err != nil {
		panic(err)
	}

	privKey, err := util.LoadKeys(keyPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client, err := eth.GetThreadsafeClient(clientURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	relay, err := eth.GetRelayHandler(common.HexToAddress(contracts.Relay), client)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	coin, err := eth.GetSwissCoin(common.HexToAddress(contracts.Swiss), client)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	node := swiss.InitSwissNode(localIP, port, globalIP, privKey, client, relay, coin)

	fmt.Println("My id is: ", node.ID)

	fmt.Println("Listening...")
	node.Start(swiss.DefaultMessageProcessor)
}
