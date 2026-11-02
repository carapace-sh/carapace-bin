package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var openai_modelsCmd = &cobra.Command{
	Use:   "models",
	Short: "List models available to you from the OpenAI API",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(openai_modelsCmd).Standalone()

	openai_modelsCmd.Flags().Bool("json", false, "Output as JSON")
	openai_modelsCmd.Flags().String("key", "", "OpenAI API key")
	openaiCmd.AddCommand(openai_modelsCmd)
}
