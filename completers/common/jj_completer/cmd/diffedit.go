package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace-jjlex/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var diffeditCmd = &cobra.Command{
	Use:   "diffedit",
	Short: "Touch up the content changes in a revision with a diff editor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffeditCmd).Standalone()

	diffeditCmd.Flags().StringP("from", "f", "", "Show changes from this revision")
	diffeditCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	diffeditCmd.Flags().Bool("restore-descendants", false, "Preserve the content (not the diff) when rebasing descendants")
	diffeditCmd.Flags().StringP("revision", "r", "", "The revision to touch up")
	diffeditCmd.Flags().StringP("to", "t", "", "Edit changes in this revision")
	diffeditCmd.Flags().String("tool", "", "Specify diff editor to be used")
	rootCmd.AddCommand(diffeditCmd)

	carapace.Gen(diffeditCmd).FlagCompletion(carapace.ActionMap{
		"from":     jj.ActionRevsets(jj.RevOpts{}.Default()),
		"revision": jj.ActionRevsets(jj.RevOpts{}.Default()),
		"to":       jj.ActionRevsets(jj.RevOpts{}.Default()),
		"tool":     bridge.ActionCarapaceBin().Split(),
	})

	carapace.Gen(diffeditCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			revisionFlag := diffeditCmd.Flag("revision")
			fromFlag := diffeditCmd.Flag("from")
			toFlag := diffeditCmd.Flag("to")

			if revisionFlag.Changed {
				return jj.ActionRevDiffs(revisionFlag.Value.String()).FilterArgs()
			}

			from, to := fromFlag.Value.String(), toFlag.Value.String()
			switch {
			case fromFlag.Changed && !toFlag.Changed:
				to = "@"
			case !fromFlag.Changed && toFlag.Changed:
				from = "@"
			case !fromFlag.Changed && !toFlag.Changed:
				return jj.ActionRevDiffs("@").FilterArgs()
			}
			return jj.ActionRevDiffs(from, to).FilterArgs()
		}),
	)
}
