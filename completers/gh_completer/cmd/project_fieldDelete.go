package cmd

import (
	"github.com/carapace-sh/carapace"
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
	project_fieldDeleteCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	project_fieldDeleteCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	project_fieldDeleteCmd.MarkFlagRequired("id")
	projectCmd.AddCommand(project_fieldDeleteCmd)

	carapace.Gen(project_fieldDeleteCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json"),
	})
}
