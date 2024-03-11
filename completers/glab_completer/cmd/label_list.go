package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var label_listCmd = &cobra.Command{
	Use:     "list [flags]",
	Short:   "List labels in repository",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(label_listCmd).Standalone()

	label_listCmd.Flags().StringP("page", "p", "", "Page number")
	label_listCmd.Flags().StringP("per-page", "P", "", "Number of items to list per page")
	labelCmd.AddCommand(label_listCmd)
}
