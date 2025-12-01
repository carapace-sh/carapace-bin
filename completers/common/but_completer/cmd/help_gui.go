package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_guiCmd = &cobra.Command{
	Use:   "gui",
	Short: "Open the GitButler GUI for the current project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_guiCmd).Standalone()

	helpCmd.AddCommand(help_guiCmd)
}
