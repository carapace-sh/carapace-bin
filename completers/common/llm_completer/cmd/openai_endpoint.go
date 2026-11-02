package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var openai_endpointCmd = &cobra.Command{
	Use:   "endpoint",
	Short: "Run against an OpenAI-compatible endpoint without logging.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(openai_endpointCmd).Standalone()

	openai_endpointCmd.Flags().StringP("attachment", "a", "", "Attachment path or URL or -")
	openai_endpointCmd.Flags().String("attachment-type", "", "\b")
	openai_endpointCmd.Flags().String("chain-limit", "", "How many chained tool responses to allow, default 5, set 0 for unlimited")
	openai_endpointCmd.Flags().Bool("chat", false, "Start an interactive chat")
	openai_endpointCmd.Flags().String("functions", "", "Python code block or file path defining functions to register as tools")
	openai_endpointCmd.Flags().StringP("header", "H", "", "Additional HTTP header")
	openai_endpointCmd.Flags().BoolP("hide-reasoning", "R", false, "Hide reasoning output")
	openai_endpointCmd.Flags().String("key", "", "API key or stored key alias to send")
	openai_endpointCmd.Flags().StringP("model", "m", "", "Model ID (required unless --models or provided by template)")
	openai_endpointCmd.Flags().Bool("models", false, "List model IDs from the endpoint and exit")
	openai_endpointCmd.Flags().Bool("no-stream", false, "Do not stream output")
	openai_endpointCmd.Flags().StringP("option", "o", "", "key/value options for the model")
	openai_endpointCmd.Flags().StringP("param", "p", "", "Parameters for template")
	openai_endpointCmd.Flags().Bool("responses", false, "Use the Responses API instead of Chat Completions")
	openai_endpointCmd.Flags().String("schema", "", "Schema DSL, JSON schema, filepath or ID")
	openai_endpointCmd.Flags().String("schema-multi", "", "JSON schema to use for multiple results")
	openai_endpointCmd.Flags().StringP("system", "s", "", "System prompt to use")
	openai_endpointCmd.Flags().StringP("template", "t", "", "Template to use")
	openai_endpointCmd.Flags().StringP("tool", "T", "", "Name of a tool to make available to the model")
	openai_endpointCmd.Flags().Bool("tools-approve", false, "Manually approve every tool execution")
	openai_endpointCmd.Flags().Bool("tools-debug", false, "Show full details of tool executions")
	openaiCmd.AddCommand(openai_endpointCmd)
}
