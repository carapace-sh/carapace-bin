package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var aiCmd = &cobra.Command{
	Use:   "ai",
	Short: "Run the AI assistant",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aiCmd).Standalone()

	aiCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(aiCmd)
}
