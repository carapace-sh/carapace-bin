package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var pane_runCmd = &cobra.Command{
	Use:   "run <pane_id> <command>",
	Short: "Run a command in a pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_runCmd).Standalone()

	paneCmd.AddCommand(pane_runCmd)

	carapace.Gen(pane_runCmd).PositionalCompletion(herdr.ActionPanes(herdr.PaneOpts{}))
}
