package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stash_help_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Save local changes to the stash and reset the working tree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_help_pushCmd).Standalone()

	stash_helpCmd.AddCommand(stash_help_pushCmd)
}
