package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var operation_showCmd = &cobra.Command{
	Use:   "show [OPTIONS] [OPERATION]",
	Short: "Show changes to the repository in an operation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operation_showCmd).Standalone()

	// Options
	operation_showCmd.Flags().Bool("no-graph", false, "Don't show the graph, show a flat list of operations")
	operation_showCmd.Flags().BoolP("patch", "p", false, "Show patch of modifications to changes")

	// Diff formatting options
	operation_showCmd.Flags().Bool("color-words", false, "Show a word-level diff with changes indicated only by color")
	operation_showCmd.Flags().Int("context", 3, "Number of lines of context to show")
	operation_showCmd.Flags().Bool("git", false, "Show a Git-format diff")
	operation_showCmd.Flags().Bool("name-only", false, "For each path, show only its path")
	operation_showCmd.Flags().StringP("revision", "r", "", "Show changes in this revision, compared to its parent(s)")
	operation_showCmd.Flags().Bool("stat", false, "Show a histogram of the changes")
	operation_showCmd.Flags().BoolP("summary", "s", false, "For each path, show only whether it was modified, added, or removed")
	operation_showCmd.Flags().String("tool", "", "Generate diff by external command")
	operation_showCmd.Flags().Bool("types", false, "For each path, show only its type before and after")

	operationCmd.AddCommand(operation_showCmd)

	carapace.Gen(operation_showCmd).PositionalCompletion(
		jj.ActionOperations(100),
	)
}
