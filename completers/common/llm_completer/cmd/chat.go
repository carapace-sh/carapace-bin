package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Hold an ongoing chat with a model.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chatCmd).Standalone()

	chatCmd.Flags().String("chain-limit", "", "How many chained tool responses to allow, default 5, set 0 for unlimited")
	chatCmd.Flags().BoolP("continue", "c", false, "Continue the most recent conversation.")
	chatCmd.Flags().String("conversation", "", "Continue the conversation with the given ID.")
	chatCmd.Flags().StringP("database", "d", "", "Path to log database")
	chatCmd.Flags().StringP("fragment", "f", "", "Fragment (alias, URL, hash or file path) to add to the prompt")
	chatCmd.Flags().String("functions", "", "Python code block or file path defining functions to register as tools")
	chatCmd.Flags().BoolP("hide-reasoning", "R", false, "Hide reasoning output")
	chatCmd.Flags().String("key", "", "API key to use")
	chatCmd.Flags().StringP("model", "m", "", "Model to use")
	chatCmd.Flags().Bool("no-stream", false, "Do not stream output")
	chatCmd.Flags().StringP("option", "o", "", "key/value options for the model")
	chatCmd.Flags().StringP("param", "p", "", "Parameters for template")
	chatCmd.Flags().StringP("system", "s", "", "System prompt to use")
	chatCmd.Flags().String("system-fragment", "", "Fragment to add to system prompt")
	chatCmd.Flags().StringP("template", "t", "", "Template to use")
	chatCmd.Flags().StringP("tool", "T", "", "Name of a tool to make available to the model")
	chatCmd.Flags().Bool("tools-approve", false, "Manually approve every tool execution")
	chatCmd.Flags().Bool("tools-debug", false, "Show full details of tool executions")
	rootCmd.AddCommand(chatCmd)
}
