package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var groupspaceCmd = &cobra.Command{
	Use:     "groupspace [-Hinp] [-o field,...] [-s field] [-S field] [-t type,...] dataset",
	Short:   "display group space usage",
	GroupID: "misc",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groupspaceCmd).Standalone()

	groupspaceCmd.Flags().BoolS("H", "H", false, "scripting mode")
	groupspaceCmd.Flags().StringArrayS("S", "S", nil, "sort descending by field")
	groupspaceCmd.Flags().BoolS("i", "i", false, "translate SID to POSIX ID")
	groupspaceCmd.Flags().BoolS("n", "n", false, "print numeric ID instead of name")
	groupspaceCmd.Flags().StringS("o", "o", "", "fields to display")
	groupspaceCmd.Flags().BoolS("p", "p", false, "display exact values")
	groupspaceCmd.Flags().StringArrayS("s", "s", nil, "sort ascending by field")
	groupspaceCmd.Flags().StringS("t", "t", "", "type filter")

	rootCmd.AddCommand(groupspaceCmd)

	carapace.Gen(groupspaceCmd).FlagCompletion(carapace.ActionMap{
		"S": carapace.ActionValues("type", "name", "used", "quota"),
		"o": carapace.ActionValues("type", "name", "used", "quota").UniqueList(","),
		"s": carapace.ActionValues("type", "name", "used", "quota"),
		"t": carapace.ActionValues("all", "posixuser", "smbuser", "posixgroup", "smbgroup").UniqueList(","),
	})

	carapace.Gen(groupspaceCmd).PositionalCompletion(
		carapace.Batch(
			zfs.ActionDatasets(zfs.DatasetOpts{Filesystem: true, Snapshot: true}),
			carapace.ActionFiles(),
		).ToA(),
	)
}
