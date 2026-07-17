package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var service_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the Tailscale Services your node can access",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_listCmd).Standalone()

	serviceCmd.AddCommand(service_listCmd)
}
