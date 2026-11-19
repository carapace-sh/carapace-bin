package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Execute a prompt",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(promptCmd).Standalone()

	promptCmd.Flags().Bool("async", false, "Run prompt asynchronously")
	promptCmd.Flags().StringP("attachment", "a", "", "Attachment path or URL or -")
	promptCmd.Flags().String("attachment-type", "", "\b")
	promptCmd.Flags().String("chain-limit", "", "How many chained tool responses to allow, default 5, set 0 for unlimited")
	promptCmd.Flags().BoolP("continue", "c", false, "Continue the most recent conversation.")
	promptCmd.Flags().String("conversation", "", "Continue the conversation with the given ID.")
	promptCmd.Flags().StringP("database", "d", "", "Path to log database")
	promptCmd.Flags().BoolP("extract", "x", false, "Extract first fenced code block")
	promptCmd.Flags().Bool("extract-last", false, "Extract last fenced code block")
	promptCmd.Flags().StringP("fragment", "f", "", "Fragment (alias, URL, hash or file path) to add to the prompt")
	promptCmd.Flags().String("functions", "", "Python code block or file path defining functions to register as tools")
	promptCmd.Flags().BoolP("hide-reasoning", "R", false, "Hide reasoning output")
	promptCmd.Flags().Bool("json", false, "Output the response as JSON, same format as llm logs --json")
	promptCmd.Flags().String("key", "", "API key to use")
	promptCmd.Flags().Bool("log", false, "Log prompt and response to the database")
	promptCmd.Flags().StringP("model", "m", "", "Model to use")
	promptCmd.Flags().BoolP("no-log", "n", false, "Don't log to database")
	promptCmd.Flags().Bool("no-stream", false, "Do not stream output")
	promptCmd.Flags().StringP("option", "o", "", "key/value options for the model")
	promptCmd.Flags().Bool("options", false, "Show options for the selected model")
	promptCmd.Flags().StringP("param", "p", "", "Parameters for template")
	promptCmd.Flags().StringP("query", "q", "", "Use first model matching these strings")
	promptCmd.Flags().String("save", "", "Save prompt with this template name")
	promptCmd.Flags().String("schema", "", "Schema DSL, JSON schema, filepath or ID")
	promptCmd.Flags().String("schema-multi", "", "Schema for multiple results")
	promptCmd.Flags().StringP("system", "s", "", "System prompt to use")
	promptCmd.Flags().String("system-fragment", "", "Fragment to add to system prompt")
	promptCmd.Flags().StringP("template", "t", "", "Template to use; can be repeated to combine templates")
	promptCmd.Flags().StringP("tool", "T", "", "Name of a tool to make available to the model")
	promptCmd.Flags().Bool("tools-approve", false, "Manually approve every tool execution")
	promptCmd.Flags().Bool("tools-debug", false, "Show full details of tool executions")
	promptCmd.Flags().BoolP("usage", "u", false, "Show token usage")
	rootCmd.AddCommand(promptCmd)
}
