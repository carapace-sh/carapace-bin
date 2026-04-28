package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_ai_lmstudioCmd = &cobra.Command{
	Use:   "lmstudio",
	Short: "Configure LM Studio as the active AI provider",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_ai_lmstudioCmd).Standalone()

	config_ai_lmstudioCmd.Flags().String("endpoint", "", "LM Studio API base endpoint (for example, http://localhost:1234/v1)")
	config_ai_lmstudioCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_ai_lmstudioCmd.Flags().String("model", "", "Preferred model name")
	config_aiCmd.AddCommand(config_ai_lmstudioCmd)

	// TODO flag completion
}
