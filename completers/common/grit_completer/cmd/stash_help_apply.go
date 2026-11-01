package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stash_help_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a stash entry without removing it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_help_applyCmd).Standalone()

	stash_helpCmd.AddCommand(stash_help_applyCmd)
}
