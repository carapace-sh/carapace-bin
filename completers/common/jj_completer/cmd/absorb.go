package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var absorbCmd = &cobra.Command{
	Use:   "absorb [OPTIONS] [PATHS]...",
	Short: "Move changes from a revision into the stack of mutable revisions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(absorbCmd).Standalone()

	absorbCmd.Flags().StringP("from", "f", "@", "Source revision to absorb from")
	absorbCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	absorbCmd.Flags().StringP("into", "t", "mutable()", "Destination revisions to absorb into")
	absorbCmd.Flags().String("to", "mutable()", "Alias for --into")
	rootCmd.AddCommand(absorbCmd)

	absorbCmd.MarkFlagsMutuallyExclusive("into", "to")

	carapace.Gen(absorbCmd).FlagCompletion(carapace.ActionMap{
		"from": jj.ActionRevs(jj.RevOption{}.Default()),
		"into": jj.ActionRevs(jj.RevOption{}.Default()),
		"to":   jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(absorbCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(carapace.Context) carapace.Action {
			return jj.ActionRevFiles(absorbCmd.Flag("from").Value.String())
		}),
	)
}
