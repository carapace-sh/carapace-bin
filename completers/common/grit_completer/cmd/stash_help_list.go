package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stash_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List stash entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_help_listCmd).Standalone()

	stash_helpCmd.AddCommand(stash_help_listCmd)
}
