package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

var localIP string
var globalIP string
var startPort int32
var endPort int32
var bootstrapIP string
var bootstrapPort int32
var keyPath string
var clientPath string
var relayContractAddress string
var coinContractAddress string

var infraNodeDeploy = &cobra.Command{
	Use:   "infra-deploy",
	Short: "Swiss cluster deploy",
	Long:  "Swisssssss clusterrrr deployyyy",
	Run: func(cmd *cobra.Command, args []string) {

		port := startPort
		for {
			if port > endPort {
				break
			}
			deployNode(
				localIP,
				globalIP,
				int(port),
				bootstrapIP,
				int(bootstrapPort),
				keyPath,
				clientPath,
				relayContractAddress,
				coinContractAddress)

			port++
		}

		fmt.Println("Deployment finished")
	},
}

func init() {

	flags := infraNodeDeploy.Flags()

	flags.StringVar(&localIP, "local-ip", "127.0.0.1", "The local ip address")
	flags.StringVar(&globalIP, "global-ip", "0.0.0.0", "The global ip address")
	flags.Int32Var(&startPort, "start-port", 8000, "The cluster deployment start port")
	flags.Int32Var(&endPort, "end-port", 8080, "The cluster deployment end port")
	flags.StringVar(&bootstrapIP, "bootstrap-ip", "127.0.0.1", "The bootstrap ip of the cluster")
	flags.Int32Var(&bootstrapPort, "bootstrap-port", 3030, "The bootstrap port of the cluster")
	flags.StringVar(&keyPath, "key", "", "The path to the private key file. For simplicity, the same private key will be held by all nodes in the local testnet")
	flags.StringVar(&clientPath, "conn", "https://ropsten.infura.io/", "The url to which the ethereum client connects to the network")
	flags.StringVar(&relayContractAddress, "relay", "", "The ethereum address of the relay handler contract")
	flags.StringVar(&coinContractAddress, "coin", "", "The ethereum address of the swiss coin contract")

	rootCmd.AddCommand(infraNodeDeploy)
}

func deployNode(localIP string, globalIP string, port int, bootstrapIP string, bootstrapPort int, keyPath string, clientURL string, relayAddr string, coinAddr string) {

	gopath := os.Getenv("GOPATH")
	command := exec.Command(fmt.Sprintf("%s/bin/bchain-routing", gopath),
		"single-infra-deploy",
		"--port", strconv.Itoa(port),
		"--bootstrap-ip", bootstrapIP,
		"--bootstrap-port", strconv.Itoa(bootstrapPort),
		"--local-ip", localIP,
		"--global-ip", globalIP,
		"--key", keyPath,
		"--conn", clientURL,
		"--relay", relayAddr,
		"--coin", coinAddr)

	fmt.Println("Deploying...")
	err := command.Start()
	if err != nil {
		fmt.Println(err)
	}
}
