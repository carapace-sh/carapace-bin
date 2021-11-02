package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var projects_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update an existing project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projects_updateCmd).Standalone()
	projects_updateCmd.Flags().String("description", "", "A description of the project")
	projects_updateCmd.Flags().String("environment", "", "The environment in which your project resides. Possible values: `Development`, `Staging`, or `Production`")
	projects_updateCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `OwnerUUID`, `OwnerID`, `Name`, `Description`, `Purpose`, `Environment`, `IsDefault`, `CreatedAt`, `UpdatedAt`")
	projects_updateCmd.Flags().Bool("is_default", false, "Set the specified project as your default project")
	projects_updateCmd.Flags().String("name", "", "A name for the project")
	projects_updateCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	projects_updateCmd.Flags().String("purpose", "", "The project's purpose")
	projectsCmd.AddCommand(projects_updateCmd)
}
