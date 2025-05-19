package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var file_showCmd = &cobra.Command{
	Use:     "show [OPTIONS] <PATH>",
	Short:   "Print contents of a file in a revision",
	Aliases: []string{},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(file_showCmd).Standalone()

	file_showCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	file_showCmd.Flags().StringP("revision", "r", "", "The revision to get the file contents from")
	fileCmd.AddCommand(file_showCmd)

	carapace.Gen(file_showCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(file_showCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return jj.ActionRevFiles(file_showCmd.Flag("revision").Value.String())
		}),
	)
}
