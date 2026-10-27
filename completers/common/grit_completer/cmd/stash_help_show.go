package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stash_help_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the summary or patch for a stash entry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stash_help_showCmd).Standalone()

	stash_helpCmd.AddCommand(stash_help_showCmd)
}
