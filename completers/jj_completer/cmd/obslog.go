package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var obslogCmd = &cobra.Command{
	Use:   "obslog",
	Short: "Show how a change has evolved",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(obslogCmd).Standalone()

	obslogCmd.Flags().Bool("color-words", false, "Show a word-level diff with changes indicated only by color")
	obslogCmd.Flags().Bool("git", false, "Show a Git-format diff")
	obslogCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	obslogCmd.Flags().StringP("limit", "l", "", "Limit number of revisions to show")
	obslogCmd.Flags().Bool("no-graph", false, "Don't show the graph, show a flat list of revisions")
	obslogCmd.Flags().BoolP("patch", "p", false, "Show patch compared to the previous version of this change")
	obslogCmd.Flags().StringP("revision", "r", "", "")
	obslogCmd.Flags().Bool("stat", false, "Show a histogram of the changes")
	obslogCmd.Flags().BoolP("summary", "s", false, "For each path, show only whether it was modified, added, or removed")
	obslogCmd.Flags().StringP("template", "T", "", "Render each revision using the given template")
	obslogCmd.Flags().String("tool", "", "Generate diff by external command")
	obslogCmd.Flags().Bool("types", false, "For each path, show only its type before and after")
	rootCmd.AddCommand(obslogCmd)
}
