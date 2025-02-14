package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var splitCmd = &cobra.Command{
	Use:   "split [OPTIONS] [PATHS]...",
	Short: "Split a revision in two",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(splitCmd).Standalone()

	splitCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	splitCmd.Flags().BoolP("interactive", "i", false, "Interactively choose which parts to split. This is the default if no paths are provided")
	splitCmd.Flags().BoolP("parallel", "p", false, "Split the revision into two parallel instead of parent and child")
	splitCmd.Flags().StringP("revision", "r", "", "The revision to split")
	rootCmd.AddCommand(splitCmd)

	carapace.Gen(splitCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(splitCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return jj.ActionRevFiles(splitCmd.Flag("revision").Value.String())
		}),
	)
}
