package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/charm"
	"github.com/spf13/cobra"
)

var fs_treeCmd = &cobra.Command{
	Use:   "tree [charm:]PATH",
	Short: "Print a file system tree from path.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fs_treeCmd).Standalone()

	fsCmd.AddCommand(fs_treeCmd)

	carapace.Gen(fs_treeCmd).PositionalCompletion(
		charm.ActionFiles(),
	)
}
