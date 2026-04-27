package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var config_ai_openrouterCmd = &cobra.Command{
	Use:   "openrouter",
	Short: "Configure OpenRouter as the active AI provider",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_ai_openrouterCmd).Standalone()

	config_ai_openrouterCmd.Flags().String("api-key", "", "OpenRouter API key. Prefer --api-key-env to avoid shell history exposure")
	config_ai_openrouterCmd.Flags().String("api-key-env", "", "Name of an environment variable holding the OpenRouter API key")
	config_ai_openrouterCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_ai_openrouterCmd.Flags().String("model", "", "Preferred model name (for example, openai/gpt-4.1-mini)")
	config_aiCmd.AddCommand(config_ai_openrouterCmd)

	// TODO flag completion
	carapace.Gen(config_ai_openrouterCmd).FlagCompletion(carapace.ActionMap{
		"api-key-env": os.ActionEnvironmentVariables(),
	})
}
