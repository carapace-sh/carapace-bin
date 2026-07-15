package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pane_renameCmd = &cobra.Command{
	Use:   "rename <pane_id> [label]",
	Short: "Rename a pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_renameCmd).Standalone()

	pane_renameCmd.Flags().Bool("clear", false, "")
	paneCmd.AddCommand(pane_renameCmd)
}
