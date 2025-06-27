package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show [OPTIONS] [REVISION]",
	Short: "Show commit description and changes in a revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()

	showCmd.Flags().Bool("color-words", false, "Show a word-level diff with changes indicated only by color")
	showCmd.Flags().String("context", "", "Number of lines of context to show")
	showCmd.Flags().Bool("git", false, "Show a Git-format diff")
	showCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	showCmd.Flags().BoolP("ignore-all-space", "w", false, "Ignore whitespace when comparing lines")
	showCmd.Flags().BoolP("ignore-space-change", "b", false, "Ignore changes in amount of whitespace when comparing lines")
	showCmd.Flags().Bool("name-only", false, "For each path, show only its path")
	showCmd.Flags().Bool("no-patch", false, "Do not show the patch")
	showCmd.Flags().Bool("stat", false, "Show a histogram of the changes")
	showCmd.Flags().BoolP("summary", "s", false, "For each path, show only whether it was modified, added, or deleted")
	showCmd.Flags().StringP("template", "T", "", "Render a revision using the given template")
	showCmd.Flags().String("tool", "", "Generate diff by external command")
	showCmd.Flags().Bool("types", false, "For each path, show only its type before and after")
	rootCmd.AddCommand(showCmd)

	carapace.Gen(showCmd).FlagCompletion(carapace.ActionMap{
		"tool": carapace.Batch(
			carapace.ActionExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	})

	carapace.Gen(showCmd).PositionalCompletion(
		jj.ActionRevs(jj.RevOption{}.Default()),
	)
}
