package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show logged prompts and their responses",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_listCmd).Standalone()

	logs_listCmd.Flags().String("conversation", "", "Show logs for this conversation ID")
	logs_listCmd.Flags().StringP("count", "n", "", "Number of entries to show - defaults to 3, use 0 for all")
	logs_listCmd.Flags().BoolP("current", "c", false, "Show logs from the current conversation")
	logs_listCmd.Flags().Bool("data", false, "Output newline-delimited JSON data for schema")
	logs_listCmd.Flags().Bool("data-array", false, "Output JSON array of data for schema")
	logs_listCmd.Flags().Bool("data-ids", false, "Attach corresponding IDs to JSON objects")
	logs_listCmd.Flags().String("data-key", "", "Return JSON objects from array in this key")
	logs_listCmd.Flags().StringP("database", "d", "", "Path to log database")
	logs_listCmd.Flags().BoolP("expand", "e", false, "Expand fragments to show their content")
	logs_listCmd.Flags().BoolP("extract", "x", false, "Extract first fenced code block")
	logs_listCmd.Flags().Bool("extract-last", false, "Extract last fenced code block")
	logs_listCmd.Flags().StringP("fragment", "f", "", "Filter for prompts using these fragments")
	logs_listCmd.Flags().String("id-gt", "", "Return responses with ID > this")
	logs_listCmd.Flags().String("id-gte", "", "Return responses with ID >= this")
	logs_listCmd.Flags().Bool("json", false, "Output logs as JSON")
	logs_listCmd.Flags().BoolP("latest", "l", false, "Return latest results matching search query")
	logs_listCmd.Flags().StringP("model", "m", "", "Filter by model or model alias")
	logs_listCmd.Flags().StringP("path", "p", "", "Path to log database")
	logs_listCmd.Flags().StringP("query", "q", "", "Search for logs matching this string")
	logs_listCmd.Flags().BoolP("response", "r", false, "Just output the last response")
	logs_listCmd.Flags().String("schema", "", "Schema DSL, JSON schema, filepath or ID")
	logs_listCmd.Flags().String("schema-multi", "", "JSON schema used for multiple results")
	logs_listCmd.Flags().BoolP("short", "s", false, "Shorter YAML output with truncated prompts")
	logs_listCmd.Flags().StringP("tool", "T", "", "Filter for prompts with results from these tools")
	logs_listCmd.Flags().Bool("tools", false, "Filter for prompts with results from any tools")
	logs_listCmd.Flags().BoolP("truncate", "t", false, "Truncate long strings in output")
	logs_listCmd.Flags().BoolP("usage", "u", false, "Include token usage")
	logsCmd.AddCommand(logs_listCmd)
}
