package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stash_help_dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Remove a stash entry without applying it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_help_dropCmd).Standalone()

	stash_helpCmd.AddCommand(stash_help_dropCmd)
}
