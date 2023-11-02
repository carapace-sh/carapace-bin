package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge work from multiple branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_mergeCmd).Standalone()

	helpCmd.AddCommand(help_mergeCmd)
}
