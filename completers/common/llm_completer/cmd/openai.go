package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var openaiCmd = &cobra.Command{
	Use:   "openai",
	Short: "Commands for working with OpenAI and OpenAI-compatible APIs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(openaiCmd).Standalone()

	rootCmd.AddCommand(openaiCmd)
}
