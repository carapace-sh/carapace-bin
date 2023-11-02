package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show commit history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logCmd).Standalone()

	logCmd.Flags().Bool("color-words", false, "Show a word-level diff with changes indicated only by color")
	logCmd.Flags().Bool("git", false, "Show a Git-format diff")
	logCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	logCmd.Flags().StringP("limit", "l", "", "Limit number of revisions to show")
	logCmd.Flags().Bool("no-graph", false, "Don't show the graph, show a flat list of revisions")
	logCmd.Flags().BoolP("patch", "p", false, "Show patch")
	logCmd.Flags().Bool("reversed", false, "Show revisions in the opposite order (older revisions first)")
	logCmd.Flags().StringSliceP("revisions", "r", []string{}, "Which revisions to show. Defaults to the `revsets.log` setting, or `@ | ancestors(immutable_heads().., 2) | heads(immutable_heads())` if it is not set")
	logCmd.Flags().Bool("stat", false, "Show a histogram of the changes")
	logCmd.Flags().BoolP("summary", "s", false, "For each path, show only whether it was modified, added, or removed")
	logCmd.Flags().StringP("template", "T", "", "Render each revision using the given template")
	logCmd.Flags().String("tool", "", "Generate diff by external command")
	logCmd.Flags().Bool("types", false, "For each path, show only its type before and after")
	rootCmd.AddCommand(logCmd)
}
