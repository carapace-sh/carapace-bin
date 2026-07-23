package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
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

	carapace.Gen(pane_renameCmd).PositionalCompletion(herdr.ActionPanes(herdr.PaneOpts{}))
}
