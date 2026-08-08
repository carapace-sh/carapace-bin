package cmd

import (
	"github.com/carapace-sh/carapace"
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
}
