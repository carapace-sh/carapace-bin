package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var netns_pidsCmd = &cobra.Command{
	Use:   "pids",
	Short: "report processes in the named network namespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(netns_pidsCmd).Standalone()

	netnsCmd.AddCommand(netns_pidsCmd)
}
