package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_project_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add the given Git repository as project for use with GitButler",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_project_addCmd).Standalone()

	help_projectCmd.AddCommand(help_project_addCmd)
}
