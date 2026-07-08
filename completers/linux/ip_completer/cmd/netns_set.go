package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var netns_setCmd = &cobra.Command{
	Use:   "set",
	Short: "assign an id to a peer network namespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(netns_setCmd).Standalone()

	netnsCmd.AddCommand(netns_setCmd)
}
