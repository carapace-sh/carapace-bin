package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var netns_monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "report as network namespace names are added or deleted",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(netns_monitorCmd).Standalone()

	netnsCmd.AddCommand(netns_monitorCmd)
}
