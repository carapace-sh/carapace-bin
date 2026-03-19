package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ai_help_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize shell integration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ai_help_initCmd).Standalone()

	ai_helpCmd.AddCommand(ai_help_initCmd)
}
