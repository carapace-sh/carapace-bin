package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var unzoneCmd = &cobra.Command{
	Use:     "unzone nsfile filesystem",
	Short:   "detach a filesystem from a user namespace",
	GroupID: "misc",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unzoneCmd).Standalone()

	rootCmd.AddCommand(unzoneCmd)

	carapace.Gen(unzoneCmd).PositionalCompletion(
		carapace.ActionFiles(),
		zfs.ActionFilesystems(),
	)
}
