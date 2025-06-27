package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var evologCmd = &cobra.Command{
	Use:     "evolog",
	Short:   "Show how a change has evolved over time",
	Aliases: []string{"obslog", "evolution-log"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evologCmd).Standalone()

	evologCmd.Flags().Bool("color-words", false, "Show a word-level diff with changes indicated only by color")
	evologCmd.Flags().String("context", "", "Number of lines of context to show")
	evologCmd.Flags().Bool("git", false, "Show a Git-format diff")
	evologCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	evologCmd.Flags().Bool("ignore-all-space", false, "Ignore whitespace when comparing lines")
	evologCmd.Flags().Bool("ignore-space-change", false, "Ignore changes in amount of whitespace when comparing lines")
	evologCmd.Flags().StringP("limit", "n", "", "Limit number of revisions to show")
	evologCmd.Flags().Bool("name-only", false, "For each path, show only its path")
	evologCmd.Flags().Bool("no-graph", false, "Don't show the graph, show a flat list of revisions")
	evologCmd.Flags().BoolP("patch", "p", false, "Show patch compared to the previous version of this change")
	evologCmd.Flags().Bool("reversed", false, "Show revisions in the opposite order (older revisions first)")
	evologCmd.Flags().StringP("revision", "r", "", "")
	evologCmd.Flags().Bool("stat", false, "Show a histogram of the changes")
	evologCmd.Flags().BoolP("summary", "s", false, "For each path, show only whether it was modified, added, or deleted")
	evologCmd.Flags().StringP("template", "T", "", "Render each revision using the given template")
	evologCmd.Flags().String("tool", "", "Generate diff by external command")
	evologCmd.Flags().Bool("types", false, "For each path, show only its type before and after")
	rootCmd.AddCommand(evologCmd)

	carapace.Gen(evologCmd).FlagCompletion(carapace.ActionMap{
		"revision": jj.ActionRevs(jj.RevOption{}.Default()),
		"tool":     bridge.ActionCarapaceBin().Split(),
	})
}
