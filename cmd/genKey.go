package cmd

import (
	"fmt"
	"os"

	"github.com/alex-d-tc/bchain-routing/util"
	"github.com/spf13/cobra"
)

var genKeyCommand = &cobra.Command{
	Use:   "gen-key",
	Short: "Generates a public/private key pair and saves it encoded in a file",
	Long:  "Generateeeeees aaaaa public/private keeeeeeeey pair and saaaaaaves it encoded in a fileeeeeeeeeeeeeeeeeeeeeeeeeeeeeee",
	Run: func(cmd *cobra.Command, args []string) {
		generateKey(genKeyPath)
	},
}

var genKeyPath string

func init() {

	flags := genKeyCommand.Flags()

	flags.StringVar(&genKeyPath, "out", "./result.key", "The output path for the generated key")

	rootCmd.AddCommand(genKeyCommand)
}

func generateKey(outPath string) {

	key, err := util.GenerateECDSAKey()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = util.WriteKeys(outPath, key)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
