package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var holdsCmd = &cobra.Command{
	Use:     "holds [-rHp] snapshot...",
	Short:   "list holds on a snapshot",
	GroupID: "misc",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(holdsCmd).Standalone()

	holdsCmd.Flags().BoolS("H", "H", false, "scripting mode")
	holdsCmd.Flags().BoolS("p", "p", false, "display timestamps as unix epoch")
	holdsCmd.Flags().BoolS("r", "r", false, "list holds on descendent snapshots")

	rootCmd.AddCommand(holdsCmd)

	carapace.Gen(holdsCmd).PositionalAnyCompletion(
		zfs.ActionSnapshots(),
	)
}
