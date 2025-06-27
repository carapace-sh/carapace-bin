package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var operation_logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show the operation log",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operation_logCmd).Standalone()

	operation_logCmd.Flags().Bool("color-words", false, "Show a word-level diff with changes indicated only by color")
	operation_logCmd.Flags().String("context", "", "Number of lines of context to show")
	operation_logCmd.Flags().Bool("git", false, "Show a Git-format diff")
	operation_logCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	operation_logCmd.Flags().Bool("ignore-all-space", false, "Ignore whitespace when comparing lines")
	operation_logCmd.Flags().Bool("ignore-space-change", false, "Ignore changes in amount of whitespace when comparing lines")
	operation_logCmd.Flags().StringP("limit", "n", "", "Limit number of operations to show")
	operation_logCmd.Flags().Bool("name-only", false, "For each path, show only its path")
	operation_logCmd.Flags().Bool("no-graph", false, "Don't show the graph, show a flat list of operations")
	operation_logCmd.Flags().BoolP("op-diff", "d", false, "Show changes to the repository at each operation")
	operation_logCmd.Flags().BoolP("patch", "p", false, "Show patch of modifications to changes (implies --op-diff)")
	operation_logCmd.Flags().Bool("reversed", false, "Show operations in the opposite order (older operations first)")
	operation_logCmd.Flags().Bool("stat", false, "Show a histogram of the changes")
	operation_logCmd.Flags().BoolP("summary", "s", false, "For each path, show only whether it was modified, added, or deleted")
	operation_logCmd.Flags().StringP("template", "T", "", "Render each operation using the given template")
	operation_logCmd.Flags().String("tool", "", "Generate diff by external command")
	operation_logCmd.Flags().Bool("types", false, "For each path, show only its type before and after")
	operationCmd.AddCommand(operation_logCmd)
}
