package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nar_packCmd = &cobra.Command{
	Use:     "pack",
	Short:   "serialise a path to stdout in NAR format",
	Aliases: []string{"dump-path"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nar_packCmd).Standalone()

	narCmd.AddCommand(nar_packCmd)

	carapace.Gen(nar_packCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
