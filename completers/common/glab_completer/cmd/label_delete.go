package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var label_deleteCmd = &cobra.Command{
	Use:   "delete [flags]",
	Short: "Delete labels for a repository or project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(label_deleteCmd).Standalone()

	labelCmd.AddCommand(label_deleteCmd)
}
