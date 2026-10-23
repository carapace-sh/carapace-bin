package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run a KCL server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverCmd).Standalone()

	serverCmd.Flags().String("http", ":2021", "set listen address")
	serverCmd.Flags().IntP("proc", "n", 4, "set max process count")
	rootCmd.AddCommand(serverCmd)
}
