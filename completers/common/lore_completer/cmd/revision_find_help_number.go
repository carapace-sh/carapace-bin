package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_find_help_numberCmd = &cobra.Command{
	Use:   "number",
	Short: "Find revision by number",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_find_help_numberCmd).Standalone()

	revision_find_helpCmd.AddCommand(revision_find_help_numberCmd)
}
