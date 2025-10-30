package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var project_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add the given Git repository as project for use with GitButler",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_addCmd).Standalone()

	project_addCmd.Flags().BoolP("help", "h", false, "Print help")
	project_addCmd.Flags().StringP("switch-to-workspace", "s", "", "The long name of the remote reference to track, like `refs/remotes/origin/main`, when switching to the workspace branch")
	projectCmd.AddCommand(project_addCmd)
}
