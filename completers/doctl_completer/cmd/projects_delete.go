package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var projects_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the specified project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projects_deleteCmd).Standalone()
	projects_deleteCmd.Flags().BoolP("force", "f", false, "Delete the project without confirmation")
	projectsCmd.AddCommand(projects_deleteCmd)
}
