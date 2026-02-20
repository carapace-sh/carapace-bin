package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_shellHookCmd = &cobra.Command{
	Use:   "shell-hook",
	Short: "Print the pixi environment activation script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_shellHookCmd).Standalone()

	helpCmd.AddCommand(help_shellHookCmd)
}
