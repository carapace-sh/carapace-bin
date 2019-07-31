package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var waitCmd = &cobra.Command{
	Use:     "wait [-Hp] [-T d|u] [-t activity,...] pool [interval]",
	Short:   "wait for background activity to stop",
	GroupID: "maintenance",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitCmd).Standalone()

	waitCmd.Flags().BoolS("H", "H", false, "scripting mode")
	waitCmd.Flags().StringS("T", "T", "", "timestamp format")
	waitCmd.Flags().BoolS("p", "p", false, "display exact values")
	waitCmd.Flags().StringS("t", "t", "", "activity types to wait for")

	rootCmd.AddCommand(waitCmd)

	carapace.Gen(waitCmd).FlagCompletion(carapace.ActionMap{
		"T": carapace.ActionValues("d", "u"),
		"t": carapace.ActionValues(
			"discard", "free", "initialize", "replace",
			"remove", "resilver", "scrub", "trim", "raidz_expand",
		).UniqueList(","),
	})

	carapace.Gen(waitCmd).PositionalCompletion(
		zfs.ActionPools(),
	)
}
