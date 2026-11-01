package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit the staged state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_commitCmd).Standalone()

	revision_helpCmd.AddCommand(revision_help_commitCmd)
}
