package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var projects_resources_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a resource by its URN",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projects_resources_getCmd).Standalone()
	projects_resources_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `URN`, `AssignedAt`, `Status`")
	projects_resources_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	projects_resourcesCmd.AddCommand(projects_resources_getCmd)
}
