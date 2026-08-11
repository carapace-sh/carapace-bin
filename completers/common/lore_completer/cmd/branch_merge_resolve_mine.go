package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_merge_resolve_mineCmd = &cobra.Command{
	Use:   "mine",
	Short: "Resolve using my changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_merge_resolve_mineCmd).Standalone()

	branch_merge_resolve_mineCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_merge_resolve_mineCmd.Flags().String("targets", "", "Path to a targets file")
	branch_merge_resolveCmd.AddCommand(branch_merge_resolve_mineCmd)

	carapace.Gen(branch_merge_resolve_mineCmd).FlagCompletion(carapace.ActionMap{
		"targets": carapace.ActionFiles(),
	})

	carapace.Gen(branch_merge_resolve_mineCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
