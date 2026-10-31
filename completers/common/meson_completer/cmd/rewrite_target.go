package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rewrite_targetCmd = &cobra.Command{
	Use:   "target",
	Short: "Modify a target",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rewrite_targetCmd).Standalone()

	rewrite_targetCmd.Flags().StringP("subdir", "s", "", "Subdirectory of the new target (only for the \"add_target\" action)")
	rewrite_targetCmd.Flags().String("type", "", "Type of the target to add (only for the \"add_target\" action)")
	applyCommonFlagCompletion(rewrite_targetCmd)
	rewriteCmd.AddCommand(rewrite_targetCmd)
}
