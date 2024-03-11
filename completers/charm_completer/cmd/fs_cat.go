package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/charm"
	"github.com/spf13/cobra"
)

var fs_catCmd = &cobra.Command{
	Use:   "cat [charm:]PATH",
	Short: "Output the content of the file at path.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fs_catCmd).Standalone()

	fsCmd.AddCommand(fs_catCmd)

	carapace.Gen(fs_catCmd).PositionalCompletion(
		charm.ActionFiles(),
	)
}
