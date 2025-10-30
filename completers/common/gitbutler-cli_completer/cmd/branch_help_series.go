package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_seriesCmd = &cobra.Command{
	Use:   "series",
	Short: "Create a new series on top of the stack",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_seriesCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_seriesCmd)
}
