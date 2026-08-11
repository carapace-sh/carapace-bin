package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Synchronize to a given state of a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_syncCmd).Standalone()

	revision_helpCmd.AddCommand(revision_help_syncCmd)
}
