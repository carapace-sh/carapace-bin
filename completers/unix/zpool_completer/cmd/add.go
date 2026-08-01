package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add [-fgLnP] [--allow-in-use --allow-replication-mismatch --allow-ashift-mismatch] [-o property=value] pool vdev...",
	Short:   "add vdevs to a pool",
	GroupID: "vdev",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().BoolS("L", "L", false, "resolve symbolic links")
	addCmd.Flags().BoolS("P", "P", false, "display full paths")
	addCmd.Flags().Bool("allow-ashift-mismatch", false, "disable ashift validation, allowing mismatched ashift values")
	addCmd.Flags().Bool("allow-in-use", false, "allow vdevs to be added even if they might be in use in another pool")
	addCmd.Flags().Bool("allow-replication-mismatch", false, "allow vdevs with conflicting replication levels to be added")
	addCmd.Flags().BoolS("f", "f", false, "force use of vdevs")
	addCmd.Flags().BoolS("g", "g", false, "display vdev GUIDs")
	addCmd.Flags().BoolS("n", "n", false, "dry-run")
	addCmd.Flags().StringArrayS("o", "o", nil, "set property")

	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).FlagCompletion(carapace.ActionMap{
		"o": zfs.ActionPoolPropertyAssignments(),
	})

	carapace.Gen(addCmd).PositionalCompletion(
		zfs.ActionPools(),
	)

	carapace.Gen(addCmd).PositionalAnyCompletion(
		carapace.Batch(
			zfs.ActionVdevTypes(),
			carapace.ActionFiles(),
		).ToA(),
	)
}
