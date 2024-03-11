package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proxyCmd = &cobra.Command{
	Use:   "proxy",
	Short: "Run local TLS proxy allowing connecting to Teleport in single-port mode.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proxyCmd).Standalone()

	rootCmd.AddCommand(proxyCmd)
}
