package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_amendCmd = &cobra.Command{
	Use:   "amend",
	Short: "Amend the latest commit's message",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_amendCmd).Standalone()

	revision_amendCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_amendCmd.Flags().Bool("stats", false, "Print stats")
	revisionCmd.AddCommand(revision_amendCmd)

	carapace.Gen(revision_amendCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
