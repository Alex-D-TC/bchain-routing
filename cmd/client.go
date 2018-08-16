package cmd

import "github.com/spf13/cobra"

var clientCmd = &cobra.Command{}

func init() {
	rootCmd.AddCommand(clientCmd)
}
