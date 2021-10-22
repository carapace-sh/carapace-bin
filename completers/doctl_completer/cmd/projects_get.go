package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var projects_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve details for a specific project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projects_getCmd).Standalone()
	projects_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `OwnerUUID`, `OwnerID`, `Name`, `Description`, `Purpose`, `Environment`, `IsDefault`, `CreatedAt`, `UpdatedAt`")
	projects_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	projectsCmd.AddCommand(projects_getCmd)
}
