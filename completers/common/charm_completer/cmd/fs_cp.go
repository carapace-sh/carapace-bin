package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/charm"
	"github.com/spf13/cobra"
)

var fs_cpCmd = &cobra.Command{
	Use:   "cp [charm:]PATH [charm:]PATH",
	Short: "Copy a file, preface source or destination with \"charm:\" to specify a remote path.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fs_cpCmd).Standalone()

	fs_cpCmd.Flags().BoolP("recursive", "r", false, "copy directories recursively")
	fsCmd.AddCommand(fs_cpCmd)

	carapace.Gen(fs_cpCmd).PositionalCompletion(
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
