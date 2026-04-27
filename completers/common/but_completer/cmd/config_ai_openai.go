package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var config_ai_openaiCmd = &cobra.Command{
	Use:   "openai",
	Short: "Configure OpenAI as the active AI provider",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_ai_openaiCmd).Standalone()

	config_ai_openaiCmd.Flags().String("api-key", "", "OpenAI API key. Prefer --api-key-env to avoid shell history exposure")
	config_ai_openaiCmd.Flags().String("api-key-env", "", "Name of an environment variable holding the OpenAI API key")
	config_ai_openaiCmd.Flags().String("endpoint", "", "Optional custom OpenAI-compatible endpoint URL")
	config_ai_openaiCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_ai_openaiCmd.Flags().String("key-option", "", "Which credential source to use")
	config_ai_openaiCmd.Flags().String("model", "", "Preferred model name (for example, gpt-5.4-nano)")
	config_aiCmd.AddCommand(config_ai_openaiCmd)

	// TODO flag completion
	carapace.Gen(config_ai_openaiCmd).FlagCompletion(carapace.ActionMap{
		"api-key-env": os.ActionEnvironmentVariables(),
	})
}
