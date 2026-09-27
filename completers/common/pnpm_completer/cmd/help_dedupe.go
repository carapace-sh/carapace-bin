package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_dedupeCmd = &cobra.Command{
	Use:   "dedupe",
	Short: "Deduplicate packages in the lockfile",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_dedupeCmd).Standalone()

	helpCmd.AddCommand(help_dedupeCmd)
}
