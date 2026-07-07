package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "View and choose prompts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(promptCmd).Standalone()

	rootCmd.AddCommand(promptCmd)
}
