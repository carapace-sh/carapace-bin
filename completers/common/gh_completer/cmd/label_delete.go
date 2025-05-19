package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var label_deleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Short: "Delete a label from a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(label_deleteCmd).Standalone()

	label_deleteCmd.Flags().Bool("confirm", false, "Confirm deletion without prompting")
	label_deleteCmd.Flags().Bool("yes", false, "Confirm deletion without prompting")
	label_deleteCmd.Flag("confirm").Hidden = true
	labelCmd.AddCommand(label_deleteCmd)

	carapace.Gen(label_deleteCmd).PositionalCompletion(
		action.ActionLabels(label_deleteCmd),
	)
}
