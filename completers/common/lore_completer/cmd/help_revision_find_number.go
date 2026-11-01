package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_find_numberCmd = &cobra.Command{
	Use:   "number",
	Short: "Find revision by number",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_find_numberCmd).Standalone()

	help_revision_findCmd.AddCommand(help_revision_find_numberCmd)
}
