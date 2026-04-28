package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var config_ai_anthropicCmd = &cobra.Command{
	Use:   "anthropic",
	Short: "Configure Anthropic as the active AI provider",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_ai_anthropicCmd).Standalone()

	config_ai_anthropicCmd.Flags().String("api-key", "", "Anthropic API key. Prefer --api-key-env to avoid shell history exposure")
	config_ai_anthropicCmd.Flags().String("api-key-env", "", "Name of an environment variable holding the Anthropic API key")
	config_ai_anthropicCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_ai_anthropicCmd.Flags().String("key-option", "", "Which credential source to use")
	config_ai_anthropicCmd.Flags().String("model", "", "Preferred model name (for example, claude-3-5-haiku-latest)")
	config_aiCmd.AddCommand(config_ai_anthropicCmd)

	carapace.Gen(config_ai_anthropicCmd).FlagCompletion(carapace.ActionMap{
		"api-key-env": os.ActionEnvironmentVariables(),
		"key-option":  carapace.ActionValues(), // TODO
		"model":       carapace.ActionValues(), // TODO model names
	})
}
