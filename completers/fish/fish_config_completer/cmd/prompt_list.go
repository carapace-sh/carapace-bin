package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var promptListCmd = &cobra.Command{
	Use:   "list",
	Short: "List available sample prompts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(promptListCmd).Standalone()

	promptCmd.AddCommand(promptListCmd)
}
