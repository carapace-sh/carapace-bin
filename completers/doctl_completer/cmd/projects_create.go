package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var projects_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projects_createCmd).Standalone()
	projects_createCmd.Flags().String("description", "", "A description of the project")
	projects_createCmd.Flags().String("environment", "", "The environment in which your project resides. Possible values: `Development`, `Staging`, or `Production`")
	projects_createCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `OwnerUUID`, `OwnerID`, `Name`, `Description`, `Purpose`, `Environment`, `IsDefault`, `CreatedAt`, `UpdatedAt`")
	projects_createCmd.Flags().String("name", "", "A name for the project (required)")
	projects_createCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	projects_createCmd.Flags().String("purpose", "", "The project's purpose (required)")
	projectsCmd.AddCommand(projects_createCmd)
}
