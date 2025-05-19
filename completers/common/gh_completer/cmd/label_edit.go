package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/color"
	"github.com/spf13/cobra"
)

var label_editCmd = &cobra.Command{
	Use:   "edit <name>",
	Short: "Edit a label",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(label_editCmd).Standalone()

	label_editCmd.Flags().StringP("color", "c", "", "Color of the label")
	label_editCmd.Flags().StringP("description", "d", "", "Description of the label")
	label_editCmd.Flags().StringP("name", "n", "", "New name of the label")
	labelCmd.AddCommand(label_editCmd)

	carapace.Gen(label_editCmd).FlagCompletion(carapace.ActionMap{
		"color": color.ActionHexColors(),
	})

	carapace.Gen(label_editCmd).PositionalCompletion(
		action.ActionLabels(label_editCmd),
	)
}
