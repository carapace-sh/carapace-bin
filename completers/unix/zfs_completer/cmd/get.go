package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get all|property[,...] [dataset...]",
	Short:   "display properties for datasets",
	GroupID: "property",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	getCmd.Flags().BoolS("H", "H", false, "scripting mode")
	getCmd.Flags().StringS("d", "d", "", "limit recursion depth")
	getCmd.Flags().BoolP("json", "j", false, "display output in JSON format")
	getCmd.Flags().Bool("json-int", false, "display numbers in integer format for JSON output")
	getCmd.Flags().StringS("o", "o", "", "columns to display")
	getCmd.Flags().BoolS("p", "p", false, "display exact values")
	getCmd.Flags().BoolS("r", "r", false, "recurse children")
	getCmd.Flags().StringS("s", "s", "", "property sources to display")
	getCmd.Flags().StringS("t", "t", "", "dataset types to display")

	rootCmd.AddCommand(getCmd)

	carapace.Gen(getCmd).FlagCompletion(carapace.ActionMap{
		"o": carapace.ActionValues("name", "property", "value", "source").UniqueList(","),
		"s": carapace.ActionValues("local", "default", "inherited", "temporary", "received", "none").UniqueList(","),
		"t": zfs.ActionDatasetTypes().UniqueList(","),
	})

	carapace.Gen(getCmd).PositionalCompletion(
		carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				carapace.ActionValues("all"),
				zfs.ActionProperties(),
				zfs.ActionReadonlyProperties(),
			).ToA().FilterArgs()
		}),
	)

	carapace.Gen(getCmd).PositionalAnyCompletion(
		zfs.ActionDatasets(zfs.DatasetOpts{Filesystem: true, Volume: true, Snapshot: true, Bookmark: true}),
	)
}
