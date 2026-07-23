package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var netns_attachCmd = &cobra.Command{
	Use:   "attach",
	Short: "create a named network namespace from a process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(netns_attachCmd).Standalone()

	netnsCmd.AddCommand(netns_attachCmd)
}
