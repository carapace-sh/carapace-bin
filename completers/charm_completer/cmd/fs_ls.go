package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/charm"
	"github.com/spf13/cobra"
)

var fs_lsCmd = &cobra.Command{
	Use:   "ls [charm:]PATH",
	Short: "List file or directory at path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fs_lsCmd).Standalone()

	fsCmd.AddCommand(fs_lsCmd)

	carapace.Gen(fs_lsCmd).PositionalCompletion(
		charm.ActionFiles(),
	)
}
