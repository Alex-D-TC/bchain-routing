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

var singleInfraNodeDeploy = &cobra.Command{
	Use:   "single-infra-deploy",
	Short: "Swiss cluster deploy",
	Long:  "Swisssssss clusterrrr deployyyy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(sLocalIP)
		fmt.Println(sGlobalIP)
		fmt.Println(sPort)
		fmt.Println(sBootstrapIP)
		fmt.Println(sBootstrapPort)
		runNode(
			sLocalIP,
			sGlobalIP,
			int(sPort),
			sBootstrapIP,
			int(sBootstrapPort),
			sKeyPath,
			sClientPath,
			sConfigPath)
	},
}

var sLocalIP string
var sGlobalIP string
var sPort int32
var sBootstrapIP string
var sBootstrapPort int32
var sKeyPath string
var sClientPath string
var sConfigPath string

func init() {
	flags := singleInfraNodeDeploy.Flags()

	flags.StringVar(&sLocalIP, "local-ip", "127.0.0.1", "The local ip address")
	flags.StringVar(&sGlobalIP, "global-ip", "0.0.0.0", "The global ip address")
	flags.Int32Var(&sPort, "port", 8000, "The cluster port")
	flags.StringVar(&sBootstrapIP, "bootstrap-ip", "127.0.0.1", "The bootstrap ip of the cluster")
	flags.Int32Var(&sBootstrapPort, "bootstrap-port", 3030, "The bootstrap port of the cluster")
	flags.StringVar(&sKeyPath, "key", "", "The path to the private key file")
	flags.StringVar(&sClientPath, "conn", "https://ropsten.infura.io/", "The url to which the ethereum client connects to the network")
	flags.StringVar(&sConfigPath, "config", "", "The config path")

	rootCmd.AddCommand(singleInfraNodeDeploy)
}

func runNode(localIP string, globalIP string, port int, bootstrapIP string, bootstrapPort int, keyPath string, clientURL string, configPath string) {

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

	err = node.JoinAndStart(swiss.DefaultMessageProcessor, bootstrapIP, int(bootstrapPort))
	if err != nil {
		fmt.Println(err)
		return
	}
}
