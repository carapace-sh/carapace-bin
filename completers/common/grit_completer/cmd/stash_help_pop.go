package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stash_help_popCmd = &cobra.Command{
	Use:   "pop",
	Short: "Apply a stash entry and remove it when there are no conflicts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_help_popCmd).Standalone()

	stash_helpCmd.AddCommand(stash_help_popCmd)
}
