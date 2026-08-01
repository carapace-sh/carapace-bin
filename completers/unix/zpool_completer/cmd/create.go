package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create [-dfn] [-o property=value]... [-O property=value]... [-m mountpoint] [-R root] [-t tname] pool vdev...",
	Short:   "create a new storage pool",
	GroupID: "creation",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().StringArrayS("O", "O", nil, "set filesystem property")
	createCmd.Flags().StringS("R", "R", "", "alternate root directory")
	createCmd.Flags().BoolS("d", "d", false, "do not enable any features")
	createCmd.Flags().BoolS("f", "f", false, "force use of vdevs")
	createCmd.Flags().StringS("m", "m", "", "mount point for root dataset")
	createCmd.Flags().BoolS("n", "n", false, "dry-run")
	createCmd.Flags().StringArrayS("o", "o", nil, "set pool property")
	createCmd.Flags().StringS("t", "t", "", "temporary pool name")

	rootCmd.AddCommand(createCmd)

	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"O": zfs.ActionPropertyAssignments(),
		"R": carapace.ActionDirectories(),
		"m": carapace.Batch(
			carapace.ActionValues("none", "legacy"),
			carapace.ActionDirectories(),
		).ToA(),
		"o": zfs.ActionPoolPropertyAssignments(),
	})

	carapace.Gen(createCmd).PositionalCompletion(
		carapace.ActionValues(),
	)

	carapace.Gen(createCmd).PositionalAnyCompletion(
		carapace.Batch(
			zfs.ActionVdevTypes(),
			carapace.ActionFiles(),
		).ToA(),
	)
}
