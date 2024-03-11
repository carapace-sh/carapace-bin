package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/time"
	"github.com/spf13/cobra"
)

var project_itemEditCmd = &cobra.Command{
	Use:   "item-edit",
	Short: "Edit an item in a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_itemEditCmd).Standalone()

	project_itemEditCmd.Flags().String("body", "", "Body of the draft issue item")
	project_itemEditCmd.Flags().Bool("clear", false, "Remove field value")
	project_itemEditCmd.Flags().String("date", "", "Date value for the field (YYYY-MM-DD)")
	project_itemEditCmd.Flags().String("field-id", "", "ID of the field to update")
	project_itemEditCmd.Flags().String("format", "", "Output format: {json}")
	project_itemEditCmd.Flags().String("id", "", "ID of the item to edit")
	project_itemEditCmd.Flags().String("iteration-id", "", "ID of the iteration value to set on the field")
	project_itemEditCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	project_itemEditCmd.Flags().String("number", "", "Number value for the field")
	project_itemEditCmd.Flags().String("project-id", "", "ID of the project to which the field belongs to")
	project_itemEditCmd.Flags().String("single-select-option-id", "", "ID of the single select option value to set on the field")
	project_itemEditCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	project_itemEditCmd.Flags().String("text", "", "Text value for the field")
	project_itemEditCmd.Flags().String("title", "", "Title of the draft issue item")
	project_itemEditCmd.MarkFlagRequired("id")
	projectCmd.AddCommand(project_itemEditCmd)

	// TODO missing flag completion
	carapace.Gen(project_itemEditCmd).FlagCompletion(carapace.ActionMap{
		"date": time.ActionDate(),
	})
}
