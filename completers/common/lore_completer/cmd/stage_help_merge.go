package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stage_help_mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Stage as a merge",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stage_help_mergeCmd).Standalone()

	stage_helpCmd.AddCommand(stage_help_mergeCmd)
}
