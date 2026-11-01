package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var promoteCmd = &cobra.Command{
	Use:     "promote clone-filesystem",
	Short:   "promote a clone to an independent filesystem",
	GroupID: "dataset",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(promoteCmd).Standalone()

	rootCmd.AddCommand(promoteCmd)

	carapace.Gen(promoteCmd).PositionalCompletion(
		zfs.ActionFilesystems(),
	)
}
