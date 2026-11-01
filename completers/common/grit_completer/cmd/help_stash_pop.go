package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_stash_popCmd = &cobra.Command{
	Use:   "pop",
	Short: "Apply a stash entry and remove it when there are no conflicts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_stash_popCmd).Standalone()

	help_stashCmd.AddCommand(help_stash_popCmd)
}
