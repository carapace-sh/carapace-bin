package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "List snapshots for mounted shares",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshotCmd).Standalone()

	snapshotCmd.Flags().BoolS("a", "a", false, "List snapshots of all mounted shares")
	snapshotCmd.Flags().StringS("f", "f", "", "Output format")
	snapshotCmd.Flags().StringS("m", "m", "", "List snapshots for the item at path")

	carapace.Gen(snapshotCmd).FlagCompletion(carapace.ActionMap{
		"f": carapace.ActionValues("Json"),
		"m": carapace.ActionDirectories(),
	})
}
