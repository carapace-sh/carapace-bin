package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_historyCmd = &cobra.Command{
	Use:   "history",
	Short: "List revisions of a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_historyCmd).Standalone()

	revision_helpCmd.AddCommand(revision_help_historyCmd)
}
