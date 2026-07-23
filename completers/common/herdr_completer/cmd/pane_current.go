package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var pane_currentCmd = &cobra.Command{
	Use:   "current",
	Short: "Show the current pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_currentCmd).Standalone()

	pane_currentCmd.Flags().Bool("current", false, "")
	pane_currentCmd.Flags().String("pane", "", "")
	paneCmd.AddCommand(pane_currentCmd)

	carapace.Gen(pane_currentCmd).FlagCompletion(carapace.ActionMap{
		"pane": herdr.ActionPanes(herdr.PaneOpts{}),
	})
}
