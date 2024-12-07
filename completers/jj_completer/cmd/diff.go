package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var diffCmd = &cobra.Command{
	Use:   "diff [OPTIONS] [PATHS]...",
	Short: "Compare file contents between two revisions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(diffCmd).Standalone()

	diffCmd.Flags().Int("context", 3, "Number of lines of context to show")
	diffCmd.Flags().Bool("color-words", false, "Show a word-level diff with changes indicated only by color")
	diffCmd.Flags().StringP("from", "-f", "", "Show changes from this revision")
	diffCmd.Flags().Bool("git", false, "Show a Git-format diff")
	diffCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	diffCmd.Flags().BoolP("ignore-all-space", "w", false, "Ignore whitespace when comparing lines")
	diffCmd.Flags().BoolP("ignore-space-change", "b", false, " Ignore changes in amount of whitespace when comparing lines")
	diffCmd.Flags().Bool("name-only", false, "For each path, show only its path")
	diffCmd.Flags().StringP("revision", "r", "", "Show changes in this revision, compared to its parent(s)")
	diffCmd.Flags().Bool("stat", false, "Show a histogram of the changes")
	diffCmd.Flags().BoolP("summary", "s", false, "For each path, show only whether it was modified, added, or removed")
	diffCmd.Flags().StringP("to", "t", "", "Show changes to this revision")
	diffCmd.Flags().String("tool", "", "Generate diff by external command")
	diffCmd.Flags().Bool("types", false, "For each path, show only its type before and after")
	rootCmd.AddCommand(diffCmd)

	diffCmd.MarkFlagsMutuallyExclusive("name-only", "summary")

	carapace.Gen(diffCmd).FlagCompletion(carapace.ActionMap{
		"from":     jj.ActionRevs(jj.RevOption{}.Default()),
		"revision": jj.ActionRevs(jj.RevOption{}.Default()),
		"to":       jj.ActionRevs(jj.RevOption{}.Default()),
		"tool":     bridge.ActionCarapaceBin().Split(),
	})

	carapace.Gen(diffCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			revisions := []string{"@"}
			if f := diffCmd.Flag("from"); f.Changed {
				revisions = []string{f.Value.String()}
			}
			if f := diffCmd.Flag("to"); f.Changed {
				revisions = append(revisions, f.Value.String())
			}
			return jj.ActionRevDiffs(revisions...)
		}),
	)
}
