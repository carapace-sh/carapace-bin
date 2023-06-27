package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var project_fieldDeleteCmd = &cobra.Command{
	Use:   "field-delete",
	Short: "Delete a field in a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_fieldDeleteCmd).Standalone()

	project_fieldDeleteCmd.Flags().String("format", "", "Output format: {json}")
	project_fieldDeleteCmd.Flags().String("id", "", "ID of the field to delete")
	projectCmd.AddCommand(project_fieldDeleteCmd)

	carapace.Gen(project_fieldDeleteCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json"),
	})
}
