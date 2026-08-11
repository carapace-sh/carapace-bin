package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_revision_bisectCmd = &cobra.Command{
	Use:   "bisect",
	Short: "Binary search for a change introduced between start (exclusive) and end (inclusive.)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_revision_bisectCmd).Standalone()

	help_revisionCmd.AddCommand(help_revision_bisectCmd)
}
