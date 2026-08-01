package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var userspaceCmd = &cobra.Command{
	Use:     "userspace [-Hinp] [-o field,...] [-s field] [-S field] [-t type,...] dataset",
	Short:   "display user space usage",
	GroupID: "misc",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(userspaceCmd).Standalone()

	userspaceCmd.Flags().BoolS("H", "H", false, "scripting mode")
	userspaceCmd.Flags().StringArrayS("S", "S", nil, "sort descending by field")
	userspaceCmd.Flags().BoolS("i", "i", false, "translate SID to POSIX ID")
	userspaceCmd.Flags().BoolS("n", "n", false, "print numeric ID instead of name")
	userspaceCmd.Flags().StringS("o", "o", "", "fields to display")
	userspaceCmd.Flags().BoolS("p", "p", false, "display exact values")
	userspaceCmd.Flags().StringArrayS("s", "s", nil, "sort ascending by field")
	userspaceCmd.Flags().StringS("t", "t", "", "type filter")

	rootCmd.AddCommand(userspaceCmd)

	carapace.Gen(userspaceCmd).FlagCompletion(carapace.ActionMap{
		"S": carapace.ActionValues("type", "name", "used", "quota"),
		"o": carapace.ActionValues("type", "name", "used", "quota").UniqueList(","),
		"s": carapace.ActionValues("type", "name", "used", "quota"),
		"t": carapace.ActionValues("all", "posixuser", "smbuser", "posixgroup", "smbgroup").UniqueList(","),
	})

	carapace.Gen(userspaceCmd).PositionalCompletion(
		carapace.Batch(
			zfs.ActionDatasets(zfs.DatasetOpts{Filesystem: true, Snapshot: true}),
			carapace.ActionFiles(),
		).ToA(),
	)
}
