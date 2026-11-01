package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_help_bisectCmd = &cobra.Command{
	Use:   "bisect",
	Short: "Binary search for a change introduced between start (exclusive) and end (inclusive.)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_help_bisectCmd).Standalone()

	revision_helpCmd.AddCommand(revision_help_bisectCmd)
}
