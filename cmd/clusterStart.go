package cmd

import (
	"fmt"

	"github.com/alex-d-tc/bchain-routing/swiss"

	"github.com/spf13/cobra"
)

var rootLocalIP string
var rootGlobalIP string
var rootPort int32

var clusterStart = &cobra.Command{
	Use:   "cluster-start",
	Short: "Swiss cluster deploy",
	Long:  "Swisssssss clusterrrr deployyyy",
	Run: func(cmd *cobra.Command, args []string) {
		runCluster(rootLocalIP, rootGlobalIP, int(rootPort))
	},
}

func init() {
	flags := clusterStart.Flags()

	flags.StringVar(&rootLocalIP, "local-ip", "127.0.0.1", "The local ip address")
	flags.StringVar(&rootGlobalIP, "global-ip", "0.0.0.0", "The global ip address")
	flags.Int32Var(&rootPort, "port", 8000, "The cluster port")

	rootCmd.AddCommand(clusterStart)
}

func runCluster(localIP string, globalIP string, port int) {

	node := swiss.InitSwissNode(localIP, port, globalIP)

	fmt.Println("Listening...")
	node.Start()
}
