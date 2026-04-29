package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_inlineCmd = &cobra.Command{
	Use:   "inline",
	Short: "Inline completion mode with small TUI overlay",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_inlineCmd).Standalone()

	ai_inlineCmd.PersistentFlags().String("api-endpoint", "", "Custom API endpoint; defaults to reading from the `ai.endpoint` setting")
	ai_inlineCmd.PersistentFlags().String("api-token", "", "Custom API token; defaults to reading from the `ai.api_token` setting")
	ai_inlineCmd.Flags().BoolP("help", "h", false, "Print help")
	ai_inlineCmd.PersistentFlags().BoolP("verbose", "v", false, "Enable verbose logging")
	aiCmd.AddCommand(ai_inlineCmd)
}
