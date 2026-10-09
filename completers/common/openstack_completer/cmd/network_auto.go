package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_autoCmd = &cobra.Command{
	Use:   "auto",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_autoCmd).Standalone()

	networkCmd.AddCommand(network_autoCmd)
}
