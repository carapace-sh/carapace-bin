package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tofu_completer/cmd/action"
	"github.com/spf13/cobra"
)

var graphCmd = &cobra.Command{
	Use:   "graph [options]",
	Short: "Generate a Graphviz graph of the steps in an operation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(graphCmd).Standalone()

	graphCmd.Flags().BoolS("compact-warnings", "compact-warnings", false, "Show warnings in a more compact form that includes only the summary messages.")
	graphCmd.Flags().BoolS("concise", "concise", false, "Disable progress-related messages.")
	graphCmd.Flags().BoolS("consolidate-errors", "consolidate-errors", false, "If OpenTofu produces any errors, attempt to consolidate similar messages into a single item.")
	graphCmd.Flags().BoolS("consolidate-warnings", "consolidate-warnings", false, "If OpenTofu produces any warnings, do not attempt to consolidate similar messages.")
	graphCmd.Flags().StringArrayS("deprecation", "deprecation", nil, "Specify what type of warnings are shown.")
	graphCmd.Flags().BoolS("draw-cycles", "draw-cycles", false, "Highlight any cycles in the graph with colored edges.")
	graphCmd.Flags().IntS("module-depth", "module-depth", -1, "(deprecated) In prior versions of OpenTofu, specified the depth of modules to graph.")
	graphCmd.Flags().BoolS("no-color", "no-color", false, "Disable virtual terminal escape sequences.")
	graphCmd.Flags().StringS("plan", "plan", "", "Render graph using the specified plan file instead of the current state.")
	graphCmd.Flags().StringS("type", "type", "", "Type of graph to output. Can be: plan, plan-refresh-only, plan-destroy, or apply.")
	graphCmd.Flags().StringArrayS("var", "var", nil, "Set a value for one of the input variables in the root module of the configuration.")
	graphCmd.Flags().StringS("var-file", "var-file", "", "Load variable values from the given file.")
	graphCmd.Flags().BoolS("verbose", "verbose", false, "Generate a graph with more information.")
	rootCmd.AddCommand(graphCmd)

	carapace.Gen(graphCmd).FlagCompletion(carapace.ActionMap{
		"deprecation": action.ActionDeprecationLevels(),
		"plan":        carapace.ActionFiles(),
		"type":        carapace.ActionValues("plan", "plan-refresh-only", "plan-destroy", "apply"),
		"var-file":    carapace.ActionFiles(),
	})
}
