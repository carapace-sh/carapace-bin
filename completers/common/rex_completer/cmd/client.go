package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "Inspect the clients connected to the server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clientCmd).Standalone()

	rootCmd.AddCommand(clientCmd)
}
