package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var interdiffCmd = &cobra.Command{
	Use:   "interdiff [OPTIONS] <--from <FROM>|--to <TO>> [PATHS]..",
	Short: "Compare the changes of two commits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(interdiffCmd).Standalone()

	interdiffCmd.Flags().Int("context", 3, "Number of lines of context to show")
	interdiffCmd.Flags().Bool("color-words", false, "Show a word-level diff with changes indicated only by color")
	interdiffCmd.Flags().String("from", "@", "Show changes from this revision")
	interdiffCmd.Flags().Bool("git", false, "Show a Git-format diff")
	interdiffCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	interdiffCmd.Flags().Bool("stat", false, "Show a histogram of the changes")
	interdiffCmd.Flags().BoolP("summary", "s", false, "For each path, show only whether it was modified, added, or removed")
	interdiffCmd.Flags().String("to", "@", "Show changes to this revision")
	interdiffCmd.Flags().String("tool", "", "Generate diff by external command")
	interdiffCmd.Flags().Bool("types", false, "For each path, show only its type before and after")
	rootCmd.AddCommand(interdiffCmd)

	carapace.Gen(interdiffCmd).FlagCompletion(carapace.ActionMap{
		"from": jj.ActionRevs(jj.RevOption{}.Default()),
		"to":   jj.ActionRevs(jj.RevOption{}.Default()),
		"tool": bridge.ActionCarapaceBin().Split(),
	})

	carapace.Gen(interdiffCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return jj.ActionRevDiffs(
				interdiffCmd.Flag("from").Value.String(),
				interdiffCmd.Flag("to").Value.String(),
			).FilterArgs()
		}),
	)
}
