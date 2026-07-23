package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var netns_list_idCmd = &cobra.Command{
	Use:   "list-id",
	Short: "list network namespace ids",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(netns_list_idCmd).Standalone()

	netnsCmd.AddCommand(netns_list_idCmd)
}
