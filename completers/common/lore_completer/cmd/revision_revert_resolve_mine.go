package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revision_revert_resolve_mineCmd = &cobra.Command{
	Use:   "mine",
	Short: "Resolve using my changes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revision_revert_resolve_mineCmd).Standalone()

	revision_revert_resolve_mineCmd.Flags().BoolP("help", "h", false, "Print help")
	revision_revert_resolve_mineCmd.Flags().String("targets", "", "Path to a targets file")
	revision_revert_resolveCmd.AddCommand(revision_revert_resolve_mineCmd)
}
