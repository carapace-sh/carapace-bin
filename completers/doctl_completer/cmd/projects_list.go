package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var projects_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List existing projects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projects_listCmd).Standalone()
	projects_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `OwnerUUID`, `OwnerID`, `Name`, `Description`, `Purpose`, `Environment`, `IsDefault`, `CreatedAt`, `UpdatedAt`")
	projects_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	projectsCmd.AddCommand(projects_listCmd)
}
