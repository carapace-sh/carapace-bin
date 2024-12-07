package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore [OPTIONS] [PATHS]...",
	Short: "Restore paths from another revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(restoreCmd).Standalone()

	restoreCmd.Flags().StringP("changes-in", "c", "", "Undo the changes in a revision as compared to the merge of its parents")
	restoreCmd.Flags().StringP("from", "f", "@", "Revision to restore from (source)")
	restoreCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	restoreCmd.Flags().Bool("restore-descendants", false, "Preserve the content (not the diff) when rebasing descendants")
	restoreCmd.Flags().StringP("to", "t", "@", "Revision to restore into (destination)")
	rootCmd.AddCommand(restoreCmd)

	carapace.Gen(restoreCmd).FlagCompletion(carapace.ActionMap{
		"from": jj.ActionRevs(jj.RevOption{}.Default()),
		"to":   jj.ActionRevs(jj.RevOption{}.Default()),
	})

	carapace.Gen(restoreCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return jj.ActionRevDiffs(
				restoreCmd.Flag("from").Value.String(),
				restoreCmd.Flag("to").Value.String(),
			).FilterArgs()
		}),
	)
}
