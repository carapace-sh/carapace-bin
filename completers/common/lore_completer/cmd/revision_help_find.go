package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_findCmd = &cobra.Command{
	Use:   "find",
	Short: "Find revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_findCmd).Standalone()

	revision_helpCmd.AddCommand(revision_help_findCmd)
}
