package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/jj"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var operation_diffCmd = &cobra.Command{
	Use:   "diff [OPTIONS]",
	Short: "Compare changes to the repository between two operations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(operation_diffCmd).Standalone()

	operation_diffCmd.Flags().Bool("color-words", false, "Show a word-level diff with changes indicated only by color")
	operation_diffCmd.Flags().String("context", "", "Number of lines of context to show")
	operation_diffCmd.Flags().StringP("from", "f", "", "Show repository changes from this operation")
	operation_diffCmd.Flags().Bool("git", false, "Show a Git-format diff")
	operation_diffCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	operation_diffCmd.Flags().Bool("ignore-all-space", false, "Ignore whitespace when comparing lines")
	operation_diffCmd.Flags().Bool("ignore-space-change", false, "Ignore changes in amount of whitespace when comparing lines")
	operation_diffCmd.Flags().Bool("name-only", false, "For each path, show only its path")
	operation_diffCmd.Flags().Bool("no-graph", false, "Don't show the graph, show a flat list of modified changes")
	operation_diffCmd.Flags().String("op", "", "Show repository changes in this operation, compared to its parent")
	operation_diffCmd.Flags().String("operation", "", "Show repository changes in this operation, compared to its parent")
	operation_diffCmd.Flags().BoolP("patch", "p", false, "Show patch of modifications to changes")
	operation_diffCmd.Flags().Bool("stat", false, "Show a histogram of the changes")
	operation_diffCmd.Flags().BoolP("summary", "s", false, "For each path, show only whether it was modified, added, or deleted")
	operation_diffCmd.Flags().StringP("to", "t", "", "Show repository changes to this operation")
	operation_diffCmd.Flags().String("tool", "", "Generate diff by external command")
	operation_diffCmd.Flags().Bool("types", false, "For each path, show only its type before and after")
	operationCmd.AddCommand(operation_diffCmd)

	operation_diffCmd.MarkFlagsMutuallyExclusive("from", "to")

	carapace.Gen(operation_diffCmd).FlagCompletion(carapace.ActionMap{
		"from":      jj.ActionRevs(jj.RevOption{}.Default()),
		"operation": jj.ActionOperations(100),
		"to":        jj.ActionRevs(jj.RevOption{}.Default()),
		"tool":      bridge.ActionCarapaceBin().Split(),
	})
}
