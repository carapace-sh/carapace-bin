package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:     "clone [-pu] [-o property=value]... snapshot filesystem|volume",
	Short:   "create a clone of a snapshot",
	GroupID: "dataset",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloneCmd).Standalone()

	cloneCmd.Flags().StringArrayS("o", "o", nil, "set property")
	cloneCmd.Flags().BoolS("p", "p", false, "create parent datasets")
	cloneCmd.Flags().BoolS("u", "u", false, "do not mount the newly created file system")

	rootCmd.AddCommand(cloneCmd)

	carapace.Gen(cloneCmd).FlagCompletion(carapace.ActionMap{
		"o": zfs.ActionPropertyAssignments(),
	})

	carapace.Gen(cloneCmd).PositionalCompletion(
		zfs.ActionSnapshots(),
	)
}
