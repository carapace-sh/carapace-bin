package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_amendCmd = &cobra.Command{
	Use:   "amend",
	Short: "Amend the latest commit's message",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_amendCmd).Standalone()

	revision_helpCmd.AddCommand(revision_help_amendCmd)
}
