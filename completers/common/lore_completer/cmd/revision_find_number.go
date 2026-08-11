package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_find_numberCmd = &cobra.Command{
	Use:   "number",
	Short: "Find revision by number",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_find_numberCmd).Standalone()

	revision_find_numberCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_findCmd.AddCommand(revision_find_numberCmd)

	carapace.Gen(revision_find_numberCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
