package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var projects_resources_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List resources assigned to a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projects_resources_listCmd).Standalone()
	projects_resources_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `URN`, `AssignedAt`, `Status`")
	projects_resources_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	projects_resourcesCmd.AddCommand(projects_resources_listCmd)
}
