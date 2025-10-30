package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var project_help_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add the given Git repository as project for use with GitButler",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_help_addCmd).Standalone()

	project_helpCmd.AddCommand(project_help_addCmd)
}
