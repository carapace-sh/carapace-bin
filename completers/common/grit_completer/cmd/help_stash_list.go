package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_stash_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List stash entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_stash_listCmd).Standalone()

	help_stashCmd.AddCommand(help_stash_listCmd)
}
