package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_find_numberCmd = &cobra.Command{
	Use:   "number",
	Short: "Find revision by number",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_find_numberCmd).Standalone()

	revision_help_findCmd.AddCommand(revision_help_find_numberCmd)
}
