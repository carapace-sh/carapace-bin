package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var netns_listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"show"},
	Short:   "list network namespaces",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(netns_listCmd).Standalone()

	netnsCmd.AddCommand(netns_listCmd)
}
