package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var label_editCmd = &cobra.Command{
	Use:   "edit [flags]",
	Short: "Edit a label in a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(label_editCmd).Standalone()

	label_editCmd.Flags().StringP("color", "c", "", "The color of the label given in 6-digit hex notation with leading ‘#’ sign.")
	label_editCmd.Flags().StringP("description", "d", "", "Label description.")
	label_editCmd.Flags().StringP("label-id", "l", "", "The label ID we are updating.")
	label_editCmd.Flags().StringP("new-name", "n", "", "The new name of the label.")
	label_editCmd.Flags().StringP("priority", "p", "", "Label priority.")
	label_editCmd.MarkFlagRequired("label-id")
	labelCmd.AddCommand(label_editCmd)
}
