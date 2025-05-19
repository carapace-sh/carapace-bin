package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/charm"
	"github.com/spf13/cobra"
)

var fs_rmCmd = &cobra.Command{
	Use:   "rm [charm:]PATH",
	Short: "Remove file or directory at path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fs_rmCmd).Standalone()

	fsCmd.AddCommand(fs_rmCmd)

	carapace.Gen(fs_rmCmd).PositionalCompletion(
		charm.ActionFiles(),
	)
}
