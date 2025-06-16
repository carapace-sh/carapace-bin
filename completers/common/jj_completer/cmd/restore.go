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
	restoreCmd.Flags().BoolP("interactive", "i", false, "Interactively choose which parts to restore")
	restoreCmd.Flags().StringP("into", "t", "@", "Revision to restore into (destination)")
	restoreCmd.Flags().Bool("restore-descendants", false, "Preserve the content (not the diff) when rebasing descendants")
	restoreCmd.Flags().String("to", "@", "Revision to restore into (alias for --into)")
	restoreCmd.Flags().String("tool", "", "Specify diff editor to be used (implies --interactive)")

	restoreCmd.MarkFlagsMutuallyExclusive("into", "to")

	rootCmd.AddCommand(restoreCmd)

	carapace.Gen(restoreCmd).FlagCompletion(carapace.ActionMap{
		"from": jj.ActionRevs(jj.RevOption{}.Default()),
		"into": jj.ActionRevs(jj.RevOption{}.Default()),
		"to":   jj.ActionRevs(jj.RevOption{}.Default()),
		"tool": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	})

	carapace.Gen(restoreCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			fromFlag := restoreCmd.Flag("from")
			intoFlag := restoreCmd.Flag("into")
			if !intoFlag.Changed {
				intoFlag = restoreCmd.Flag("to")
			}

			if !fromFlag.Changed && !intoFlag.Changed {
				fromFlag.Value.Set("parents(@)")
				intoFlag.Value.Set("@")
			}

			return jj.ActionRevDiffs(
				fromFlag.Value.String(),
				intoFlag.Value.String(),
			).FilterArgs()
		}),
	)
}
