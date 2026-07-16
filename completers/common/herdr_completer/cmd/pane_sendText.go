package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var pane_sendTextCmd = &cobra.Command{
	Use:   "send-text <pane_id> <text>",
	Short: "Send literal text to a pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_sendTextCmd).Standalone()

	paneCmd.AddCommand(pane_sendTextCmd)

	carapace.Gen(pane_sendTextCmd).PositionalCompletion(herdr.ActionPanes(herdr.PaneOpts{}))
}
