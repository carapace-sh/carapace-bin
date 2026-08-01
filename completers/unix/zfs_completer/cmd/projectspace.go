package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var projectspaceCmd = &cobra.Command{
	Use:     "projectspace [-Hp] [-o field,...] [-s field] [-S field] dataset",
	Short:   "display project space usage",
	GroupID: "misc",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projectspaceCmd).Standalone()

	projectspaceCmd.Flags().BoolS("H", "H", false, "scripting mode")
	projectspaceCmd.Flags().StringArrayS("S", "S", nil, "sort descending by field")
	projectspaceCmd.Flags().StringS("o", "o", "", "fields to display")
	projectspaceCmd.Flags().BoolS("p", "p", false, "display exact values")
	projectspaceCmd.Flags().StringArrayS("s", "s", nil, "sort ascending by field")

	rootCmd.AddCommand(projectspaceCmd)

	carapace.Gen(projectspaceCmd).FlagCompletion(carapace.ActionMap{
		"S": carapace.ActionValues("type", "name", "used", "quota"),
		"o": carapace.ActionValues("type", "name", "used", "quota").UniqueList(","),
		"s": carapace.ActionValues("type", "name", "used", "quota"),
	})

	carapace.Gen(projectspaceCmd).PositionalCompletion(
		carapace.Batch(
			zfs.ActionDatasets(zfs.DatasetOpts{Filesystem: true, Snapshot: true}),
			carapace.ActionFiles(),
		).ToA(),
	)
}
