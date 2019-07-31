package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var zoneCmd = &cobra.Command{
	Use:     "zone nsfile filesystem",
	Short:   "attach a filesystem to a user namespace",
	GroupID: "misc",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(zoneCmd).Standalone()

	rootCmd.AddCommand(zoneCmd)

	carapace.Gen(zoneCmd).PositionalCompletion(
		carapace.ActionFiles(),
		zfs.ActionFilesystems(),
	)
}
