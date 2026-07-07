package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var promptShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show sample prompts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(promptShowCmd).Standalone()

	promptCmd.AddCommand(promptShowCmd)
}
