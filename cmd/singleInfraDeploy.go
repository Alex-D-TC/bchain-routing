package cmd

import (
	"fmt"

	"github.com/alex-d-tc/bchain-routing/swiss"
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
		runNode(sLocalIP, sGlobalIP, int(sPort), sBootstrapIP, int(sBootstrapPort))
	},
}

var sLocalIP string
var sGlobalIP string
var sPort int32
var sBootstrapIP string
var sBootstrapPort int32

func init() {
	flags := singleInfraNodeDeploy.Flags()

	flags.StringVar(&sLocalIP, "local-ip", "127.0.0.1", "The local ip address")
	flags.StringVar(&sGlobalIP, "global-ip", "0.0.0.0", "The global ip address")
	flags.Int32Var(&sPort, "port", 8000, "The cluster port")
	flags.StringVar(&sBootstrapIP, "bootstrap-ip", "127.0.0.1", "The bootstrap ip of the cluster")
	flags.Int32Var(&sBootstrapPort, "bootstrap-port", 3030, "The bootstrap port of the cluster")

	rootCmd.AddCommand(singleInfraNodeDeploy)
}

func runNode(localIP string, globalIP string, port int, bootstrapIP string, bootstrapPort int) {

	node := swiss.InitSwissNode(localIP, port, globalIP)

	err := node.JoinAndStart(bootstrapIP, int(bootstrapPort))
	if err != nil {
		fmt.Println(err)
		return
	}
}
