package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
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
	absorbCmd.Flags().StringSliceP("into", "t", []string{"mutable()"}, "Destination revisions to absorb into")
	absorbCmd.Flags().StringSlice("to", []string{"mutable()"}, "Destination revisions to absorb into")
	rootCmd.AddCommand(absorbCmd)

	absorbCmd.MarkFlagsMutuallyExclusive("into", "to")

	carapace.Gen(absorbCmd).FlagCompletion(carapace.ActionMap{
		"from": jj.ActionRevsets(jj.RevOpts{}.Default()),
		"into": jj.ActionRevsets(jj.RevOpts{}.Default()),
		"to":   jj.ActionRevsets(jj.RevOpts{}.Default()),
	})

	carapace.Gen(absorbCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(carapace.Context) carapace.Action {
			return jj.ActionRevFiles(absorbCmd.Flag("from").Value.String())
		}),
	)
}
