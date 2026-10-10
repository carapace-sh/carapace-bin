package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var label_deleteCmd = &cobra.Command{
	Use:   "delete <name> [flags]",
	Short: "Delete a label from a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(label_deleteCmd).Standalone()

	labelCmd.AddCommand(label_deleteCmd)

	carapace.Gen(label_deleteCmd).PositionalCompletion(
		action.ActionLabels(label_deleteCmd),
	)
}
