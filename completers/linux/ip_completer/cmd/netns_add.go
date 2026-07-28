package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var netns_addCmd = &cobra.Command{
	Use:   "add",
	Short: "create a new named network namespace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(netns_addCmd).Standalone()

	netnsCmd.AddCommand(netns_addCmd)
}
