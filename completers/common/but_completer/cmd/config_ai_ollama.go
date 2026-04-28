package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_ai_ollamaCmd = &cobra.Command{
	Use:   "ollama",
	Short: "Configure Ollama as the active AI provider",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_ai_ollamaCmd).Standalone()

	config_ai_ollamaCmd.Flags().String("endpoint", "", "Ollama endpoint in host:port form (for example, localhost:11434)")
	config_ai_ollamaCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_ai_ollamaCmd.Flags().String("model", "", "Preferred model name")
	config_aiCmd.AddCommand(config_ai_ollamaCmd)

	// TODO flag completion
}
