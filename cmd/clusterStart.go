package cmd

import (
	"fmt"

	"github.com/alex-d-tc/bchain-routing/routingdriver"
	"github.com/alex-d-tc/bchain-routing/util"
	"github.com/spf13/cobra"
	"secondbit.org/wendy"
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
	node := wendy.NewNode(util.NodeIDFromStringSHA(localIP+" "+string(port)), localIP, globalIP, "", int(port))
	cluster := wendy.NewCluster(node, routingdriver.Credentials{})

	fmt.Println("Listening...")
	cluster.Listen()
}
