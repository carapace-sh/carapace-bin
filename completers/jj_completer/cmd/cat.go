package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var catCmd = &cobra.Command{
	Use:     "cat [OPTIONS] <PATH>",
	Short:   "Print contents of a file in a revision",
	Aliases: []string{"print"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(catCmd).Standalone()

	catCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	catCmd.Flags().StringP("revision", "r", "", "The revision to get the file contents from")
	rootCmd.AddCommand(catCmd)

	carapace.Gen(catCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevs(jj.RevsOption{}.Default()),
	})

	carapace.Gen(catCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return jj.ActionRevFiles(catCmd.Flag("revision").Value.String())
		}),
	)
}
