package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var file_annotateCmd = &cobra.Command{
	Use:   "annotate [OPTIONS] <PATHS>...",
	Short: "Show the source change for each line of the target file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_annotateCmd).Standalone()

	file_annotateCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	file_annotateCmd.Flags().StringP("revision", "r", "@", "Revision to start at")
	fileCmd.AddCommand(file_annotateCmd)

	carapace.Gen(file_annotateCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevs(jj.RevOption{}.Default()).UniqueList(","),
	})

	carapace.Gen(file_annotateCmd).PositionalAnyCompletion(
		carapace.ActionFiles().FilterArgs(),
	)
}
