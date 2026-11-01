package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_stash_dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Remove a stash entry without applying it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_stash_dropCmd).Standalone()

	help_stashCmd.AddCommand(help_stash_dropCmd)
}
