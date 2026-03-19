package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_ai_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize shell integration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_ai_initCmd).Standalone()

	help_aiCmd.AddCommand(help_ai_initCmd)
}
