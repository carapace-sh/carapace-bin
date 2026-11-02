package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rewrite_tgtCmd = &cobra.Command{
	Use:   "tgt",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rewrite_tgtCmd).Standalone()

	rewrite_tgtCmd.Flags().StringP("subdir", "s", "", "Subdirectory of the new target (only for the \"add_target\" action)")
	rewrite_tgtCmd.Flags().String("type", "", "Type of the target to add (only for the \"add_target\" action)")
	applyCommonFlagCompletion(rewrite_tgtCmd)
	rewriteCmd.AddCommand(rewrite_tgtCmd)
}
