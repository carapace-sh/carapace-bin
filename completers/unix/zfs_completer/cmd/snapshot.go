package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var snapshotCmd = &cobra.Command{
	Use:     "snapshot [-r] [-o property=value]... dataset@snapname...",
	Short:   "create a snapshot",
	Aliases: []string{"snap"},
	GroupID: "dataset",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snapshotCmd).Standalone()

	snapshotCmd.Flags().StringArrayS("o", "o", nil, "set property")
	snapshotCmd.Flags().BoolS("r", "r", false, "recursively snapshot children")

	rootCmd.AddCommand(snapshotCmd)

	carapace.Gen(snapshotCmd).FlagCompletion(carapace.ActionMap{
		"o": zfs.ActionPropertyAssignments(),
	})

	carapace.Gen(snapshotCmd).PositionalAnyCompletion(
		carapace.ActionMultiPartsN("@", 2, func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return zfs.ActionFilesystems().Suffix("@")
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
