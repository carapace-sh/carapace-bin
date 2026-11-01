package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var receiveCmd = &cobra.Command{
	Use:     "receive [-FMdehnsuvx] [-o property=value]... [-x property]... filesystem|volume|snapshot",
	Short:   "receive a send stream",
	Aliases: []string{"recv"},
	GroupID: "send",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(receiveCmd).Standalone()

	receiveCmd.Flags().BoolS("A", "A", false, "abort an interrupted receive")
	receiveCmd.Flags().BoolS("F", "F", false, "force rollback before receive")
	receiveCmd.Flags().BoolS("M", "M", false, "force unmount during receive")
	receiveCmd.Flags().BoolS("c", "c", false, "repair corruption from stream")
	receiveCmd.Flags().BoolS("d", "d", false, "discard pool name from send path")
	receiveCmd.Flags().BoolS("e", "e", false, "use only final path component")
	receiveCmd.Flags().BoolS("h", "h", false, "skip receiving holds")
	receiveCmd.Flags().BoolS("n", "n", false, "dry-run")
	receiveCmd.Flags().StringArrayS("o", "o", nil, "set property")
	receiveCmd.Flags().BoolS("s", "s", false, "save partial state on interrupt")
	receiveCmd.Flags().BoolS("u", "u", false, "do not mount received filesystem")
	receiveCmd.Flags().BoolS("v", "v", false, "verbose")
	receiveCmd.Flags().StringArrayS("x", "x", nil, "exclude property")

	rootCmd.AddCommand(receiveCmd)

	carapace.Gen(receiveCmd).FlagCompletion(carapace.ActionMap{
		"o": zfs.ActionPropertyAssignments(),
		"x": zfs.ActionProperties(),
	})

	carapace.Gen(receiveCmd).PositionalCompletion(
		zfs.ActionDatasets(zfs.DatasetOpts{Filesystem: true, Volume: true, Snapshot: true}),
	)
}
