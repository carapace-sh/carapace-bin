package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var branch_merge_resolve_theirsCmd = &cobra.Command{
	Use:   "theirs",
	Short: "Resolve using their changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branch_merge_resolve_theirsCmd).Standalone()

	branch_merge_resolve_theirsCmd.Flags().BoolP("help", "h", false, "Print help")
	branch_merge_resolve_theirsCmd.Flags().String("targets", "", "Path to a targets file")
	branch_merge_resolveCmd.AddCommand(branch_merge_resolve_theirsCmd)

	carapace.Gen(branch_merge_resolve_theirsCmd).FlagCompletion(carapace.ActionMap{
		"targets": carapace.ActionFiles(),
	})

	carapace.Gen(branch_merge_resolve_theirsCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
