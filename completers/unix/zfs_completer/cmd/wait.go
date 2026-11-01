package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var waitCmd = &cobra.Command{
	Use:     "wait [-t activity,...] dataset",
	Short:   "wait for background activity to stop",
	GroupID: "misc",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitCmd).Standalone()

	waitCmd.Flags().StringS("t", "t", "", "activity type to wait for")

	rootCmd.AddCommand(waitCmd)

	carapace.Gen(waitCmd).FlagCompletion(carapace.ActionMap{
		"t": carapace.ActionValues("deleteq").UniqueList(","),
	})

	carapace.Gen(waitCmd).PositionalCompletion(
		zfs.ActionFilesystems(),
	)
}
