package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_evalCmd = &cobra.Command{
	Use:   "eval",
	Short: "Run eval cases…",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_evalCmd).Standalone()

	plugin_evalCmd.Flags().String("ablation", "", "Run a no-plugin baseline arm and report the score delta…")
	plugin_evalCmd.Flags().Bool("allow-real-servers", false, "With --mocks record: also start the plugin's REAL MCP server processes for servers that have no…")
	plugin_evalCmd.Flags().StringSlice("allow-tools", nil, "Operator grant for gated tools (Bash, Write, Edit, WebFetch, mcp__*)")
	plugin_evalCmd.Flags().String("case", "", "Filter cases by name glob")
	plugin_evalCmd.Flags().StringP("concurrency", "j", "", "Run up to <n> agent runs at once (1-8; default 1)")
	plugin_evalCmd.Flags().StringSlice("disallow-tools", nil, "Built-in tools to keep from Claude in every case, even if granted (e.g. Agent TaskStop)")
	plugin_evalCmd.Flags().String("eval-dir", "", "Directory name (below the plugin) that holds the eval cases; results go to <plugin>/<dir>/results/…")
	plugin_evalCmd.Flags().BoolP("help", "h", false, "Display help for command")
	plugin_evalCmd.Flags().String("json", "", "Print the full run result (prompts, graders, per-run scores) as JSON to stdout, or write it to…")
	plugin_evalCmd.Flags().String("judge-model", "", "Override LLM-grader model")
	plugin_evalCmd.Flags().Bool("keep-temp", false, "Preserve scaffold dirs for debugging")
	plugin_evalCmd.Flags().String("max-cost-usd", "", "Optional hard cost ceiling; abort and report partial results if hit (exit 2)")
	plugin_evalCmd.Flags().String("mocks", "", "Mock stand-ins for MCP servers, from <eval dir>/mocks/ (record | off; default: record)")
	plugin_evalCmd.Flags().String("model", "", "Override model for all cases")
	plugin_evalCmd.Flags().Bool("no-publish", false, "Keep the HTML report local only; skip publishing it to claude.ai")
	plugin_evalCmd.Flags().Bool("no-scaffold", false, "Explicitly skip scaffold_script")
	plugin_evalCmd.Flags().String("output-dir", "", "Directory for aggregate-result.json")
	plugin_evalCmd.Flags().Bool("publish-report", false, "Also require publishing the report to claude.ai (already the default when your account supports…")
	plugin_evalCmd.Flags().String("report", "", "Write the self-contained HTML report (scores, prompts, grader verdicts) to <path> instead of the…")
	plugin_evalCmd.Flags().String("runs", "", "Override per-case runs")
	plugin_evalCmd.Flags().Bool("scaffold", false, "Run each case's scaffold_script…")
	plugin_evalCmd.Flags().StringSlice("tag", nil, "Filter cases by tag (repeatable)")
	plugin_evalCmd.Flags().String("threshold", "", "Exit 1 if any case score is below this threshold")
	plugin_evalCmd.Flags().Bool("trust-plugin", false, "Assert that you trust this plugin's code and eval suite, and skip the first-run trust prompt…")
	plugin_evalCmd.Flags().Bool("verbose", false, "Log per-message trace events to the debug log (use --debug-file to read them)")
	plugin_evalCmd.Flag("json").NoOptDefVal = " "
	pluginCmd.AddCommand(plugin_evalCmd)

	carapace.Gen(plugin_evalCmd).FlagCompletion(carapace.ActionMap{
		"eval-dir": carapace.ActionDirectories(),
		"model":    carapace.ActionValues("fable", "haiku", "opus", "sonnet"),
	})
}
