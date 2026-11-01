package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list [dataset...]",
	Short:   "list datasets and their properties",
	GroupID: "property",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().BoolS("H", "H", false, "scripting mode")
	listCmd.Flags().StringArrayS("S", "S", nil, "sort descending by property")
	listCmd.Flags().StringS("d", "d", "", "limit recursion depth")
	listCmd.Flags().BoolP("json", "j", false, "JSON output")
	listCmd.Flags().Bool("json-int", false, "print numbers as integers in JSON output")
	listCmd.Flags().StringS("o", "o", "", "properties to display")
	listCmd.Flags().BoolS("p", "p", false, "display exact values")
	listCmd.Flags().BoolS("r", "r", false, "recurse children")
	listCmd.Flags().StringArrayS("s", "s", nil, "sort ascending by property")
	listCmd.Flags().StringS("t", "t", "", "dataset types to display")

	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"S": zfs.ActionSortProperties(),
		"o": zfs.ActionListColumns().UniqueList(","),
		"s": zfs.ActionSortProperties(),
		"t": zfs.ActionDatasetTypes().UniqueList(","),
	})

	carapace.Gen(listCmd).PositionalAnyCompletion(
		zfs.ActionDatasets(zfs.DatasetOpts{Filesystem: true, Volume: true, Snapshot: true}),
	)
}
