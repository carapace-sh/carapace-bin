package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/charm"
	"github.com/spf13/cobra"
)

var fs_mvCmd = &cobra.Command{
	Use:   "mv [charm:]PATH [charm:]PATH",
	Short: "Move a file, preface source or destination with \"charm:\" to specify a remote path.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fs_mvCmd).Standalone()

	fs_mvCmd.Flags().BoolP("recursive", "r", false, "move directories recursively")
	fsCmd.AddCommand(fs_mvCmd)

	carapace.Gen(fs_mvCmd).PositionalCompletion(
		carapace.Batch(
			charm.ActionFiles(),
			carapace.ActionFiles(),
		).ToA(),
		carapace.Batch(
			charm.ActionFiles(),
			carapace.ActionFiles(),
		).ToA(),
	)

}
