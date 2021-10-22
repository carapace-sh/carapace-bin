package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var projects_resources_assignCmd = &cobra.Command{
	Use:   "assign",
	Short: "Assign one or more resources to a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projects_resources_assignCmd).Standalone()
	projects_resources_assignCmd.Flags().StringSlice("resource", []string{}, "URNs specifying resources to assign to the project")
	projects_resourcesCmd.AddCommand(projects_resources_assignCmd)
}
