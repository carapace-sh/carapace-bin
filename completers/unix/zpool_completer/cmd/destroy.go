package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var destroyCmd = &cobra.Command{
	Use:     "destroy [-f] pool",
	Short:   "destroy a storage pool",
	GroupID: "creation",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(destroyCmd).Standalone()

	destroyCmd.Flags().BoolS("f", "f", false, "force destroy")

	rootCmd.AddCommand(destroyCmd)

	carapace.Gen(destroyCmd).PositionalCompletion(
		zfs.ActionPools(),
	)
}
