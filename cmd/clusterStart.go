package cmd

import (
	"fmt"
	"os"

	"github.com/alex-d-tc/bchain-routing/swiss"
	"github.com/alex-d-tc/bchain-routing/util"

	"github.com/spf13/cobra"
)

var rootLocalIP string
var rootGlobalIP string
var rootPort int32
var rootKeyPath string

var clusterStart = &cobra.Command{
	Use:   "cluster-start",
	Short: "Swiss cluster deploy",
	Long:  "Swisssssss clusterrrr deployyyy",
	Run: func(cmd *cobra.Command, args []string) {
		runCluster(rootLocalIP, rootGlobalIP, int(rootPort), rootKeyPath)
	},
}

func init() {
	flags := clusterStart.Flags()

	flags.StringVar(&rootLocalIP, "local-ip", "127.0.0.1", "The local ip address")
	flags.StringVar(&rootGlobalIP, "global-ip", "0.0.0.0", "The global ip address")
	flags.Int32Var(&rootPort, "port", 8000, "The cluster port")
	flags.StringVar(&rootKeyPath, "key", "", "The path to the private key file")

	rootCmd.AddCommand(clusterStart)
}

func runCluster(localIP string, globalIP string, port int, keyPath string) {

	privKey, err := util.LoadKeys(keyPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	node := swiss.InitSwissNode(localIP, port, globalIP, privKey)

	fmt.Println("My id is: ", node.ID)

	fmt.Println("Listening...")
	node.Start(swiss.DefaultMessageProcessor)
}
