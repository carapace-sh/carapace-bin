package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Undo fetches back to the specified SVN revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resetCmd).Standalone()

	resetCmd.Flags().BoolP("parent", "p", false, "Discard the specified revision, keeping nearest parent")
	resetCmd.Flags().StringP("revision", "r", "", "Most recent revision to keep")
	rootCmd.AddCommand(resetCmd)

	carapace.Gen(resetCmd).FlagCompletion(carapace.ActionMap{
		"revision": carapace.ActionValues(),
	})
}
