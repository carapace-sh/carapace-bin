package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pane_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List panes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_listCmd).Standalone()

	pane_listCmd.Flags().String("workspace", "", "")
	paneCmd.AddCommand(pane_listCmd)
}
