package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var unloadKeyCmd = &cobra.Command{
	Use:     "unload-key [-r] [-a] [dataset]",
	Short:   "unload encryption key for a dataset",
	GroupID: "encryption",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unloadKeyCmd).Standalone()

	unloadKeyCmd.Flags().BoolS("a", "a", false, "unload keys for all encryption roots in all imported pools")
	unloadKeyCmd.Flags().BoolS("r", "r", false, "unload keys recursively")

	rootCmd.AddCommand(unloadKeyCmd)

	carapace.Gen(unloadKeyCmd).PositionalCompletion(
		zfs.ActionFilesystems(),
	)
}
