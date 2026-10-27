package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_stash_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Save local changes to the stash and reset the working tree",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_stash_pushCmd).Standalone()

	help_stashCmd.AddCommand(help_stash_pushCmd)
}
