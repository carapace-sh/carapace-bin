package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/color"
	"github.com/spf13/cobra"
)

var label_createCmd = &cobra.Command{
	Use:     "create [flags]",
	Short:   "Create labels for a repository or project.",
	Aliases: []string{"new"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(label_createCmd).Standalone()

	label_createCmd.Flags().StringP("color", "c", "", "Color of the label, in plain or HEX code.")
	label_createCmd.Flags().StringP("description", "d", "", "Label description.")
	label_createCmd.Flags().StringP("name", "n", "", "Name of the label.")
	label_createCmd.MarkFlagRequired("name")
	labelCmd.AddCommand(label_createCmd)

	carapace.Gen(label_createCmd).FlagCompletion(carapace.ActionMap{
		"color": color.ActionHexColors(),
		"name":  action.ActionLabels(label_createCmd),
	})
}
