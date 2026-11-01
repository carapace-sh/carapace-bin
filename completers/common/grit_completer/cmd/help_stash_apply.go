package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_stash_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a stash entry without removing it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_stash_applyCmd).Standalone()

	help_stashCmd.AddCommand(help_stash_applyCmd)
}
