package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_ai_initCmd = &cobra.Command{
	Use:    "init",
	Short:  "This command is no longer necessary. If you have it in your shell init file, feel free to remove it",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_ai_initCmd).Standalone()

	help_aiCmd.AddCommand(help_ai_initCmd)
}
