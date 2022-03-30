package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var label_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List labels in a repository",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(label_listCmd).Standalone()
	label_listCmd.Flags().IntP("limit", "L", 30, "Maximum number of items to fetch")
	label_listCmd.Flags().BoolP("web", "w", false, "List labels in the web browser")
	labelCmd.AddCommand(label_listCmd)
}
