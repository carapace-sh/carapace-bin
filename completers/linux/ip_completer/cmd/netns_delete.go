package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var netns_deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "delete named network namespace(s)",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(netns_deleteCmd).Standalone()

	netnsCmd.AddCommand(netns_deleteCmd)
}
