package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var duo_askCmd = &cobra.Command{
	Use:   "ask <prompt>",
	Short: "Generate Git commands from natural language.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(duo_askCmd).Standalone()

	duo_askCmd.Flags().Bool("git", false, "Ask a question about Git.")
	duoCmd.AddCommand(duo_askCmd)
}
