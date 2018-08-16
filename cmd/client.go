package cmd

import (
	"fmt"

	"github.com/alex-d-tc/bchain-routing/swiss"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client-start",
	Short: "Start a client node",
	Long:  "Starrrrt a cliennnnt nooooode",
	Run: func(cmd *cobra.Command, args []string) {
		runClient(*clientLocalIPP, *clientGlobalIPP, int(*clientPortP), *clientBootstrapIPP, int(*clientBootstrapPortP))
	},
}

var clientPortP *int32
var clientLocalIPP *string
var clientGlobalIPP *string
var clientBootstrapIPP *string
var clientBootstrapPortP *int32

func init() {
	flags := clientCmd.Flags()

	flags.Int32Var(clientPortP, "port", 8080, "The client port")
	flags.StringVar(clientLocalIPP, "local-ip", "127.0.0.1", "The local ip")
	flags.StringVar(clientGlobalIPP, "global-ip", "0.0.0.0", "The global ip")
	flags.StringVar(clientBootstrapIPP, "bootstrap-ip", "127.0.0.1", "The bootstrap node ip")
	flags.Int32Var(clientBootstrapPortP, "bootstrap-port", 8000, "The client bootstrap port")

	rootCmd.AddCommand(clientCmd)
}

func runClient(localIP string, publicIP string, port int, bootstrapIP string, bootstrapPort int) {
	node := swiss.InitSwissNode(localIP, port, publicIP)
	err := node.JoinAndStart(bootstrapIP, bootstrapPort)
	if err != nil {
		fmt.Println(err)
	}
}
