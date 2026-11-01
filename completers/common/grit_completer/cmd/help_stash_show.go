package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_stash_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the summary or patch for a stash entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_stash_showCmd).Standalone()

	help_stashCmd.AddCommand(help_stash_showCmd)
}
