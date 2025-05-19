package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var nar_dumpPathCmd = &cobra.Command{
	Use:   "dump-path",
	Short: "serialise a path to stdout in NAR format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(nar_dumpPathCmd).Standalone()

	narCmd.AddCommand(nar_dumpPathCmd)

	carapace.Gen(nar_dumpPathCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
