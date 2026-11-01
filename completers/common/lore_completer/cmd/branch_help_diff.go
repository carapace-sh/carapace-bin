package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_help_diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Diff two branches using the common ancestor base revision Will calculate the set of changes between source branch latest revision and the base revision that is not in the set of changes between the target branch latest revision and the base revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_help_diffCmd).Standalone()

	branch_helpCmd.AddCommand(branch_help_diffCmd)
}
