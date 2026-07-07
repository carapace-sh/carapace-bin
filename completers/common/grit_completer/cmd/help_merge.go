package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge another branch into the current one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_mergeCmd).Standalone()

	helpCmd.AddCommand(help_mergeCmd)
}
