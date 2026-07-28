package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var netns_identifyCmd = &cobra.Command{
	Use:   "identify",
	Short: "report network namespace names for a process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(netns_identifyCmd).Standalone()

	netnsCmd.AddCommand(netns_identifyCmd)
}
