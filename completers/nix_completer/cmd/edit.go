package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:     "edit",
	Short:   "open the Nix expression of a Nix package in $EDITOR",
	GroupID: "infrequently used",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(editCmd).Standalone()

	rootCmd.AddCommand(editCmd)

	addEvaluationFlags(editCmd)
	addFlakeFlags(editCmd)
	addLoggingFlags(editCmd)

	// TODO positional completion
}
