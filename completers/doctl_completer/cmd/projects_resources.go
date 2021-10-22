package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var projects_resourcesCmd = &cobra.Command{
	Use:   "resources",
	Short: "Manage resources assigned to a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projects_resourcesCmd).Standalone()
	projectsCmd.AddCommand(projects_resourcesCmd)
}
