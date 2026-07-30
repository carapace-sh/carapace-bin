package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var renewCmd = &cobra.Command{
	Use:   "renew",
	Short: "renew the IP address for an adapter",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(renewCmd).Standalone()
	rootCmd.AddCommand(renewCmd)
}
