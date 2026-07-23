package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/herdr"
	"github.com/spf13/cobra"
)

var pane_sendKeysCmd = &cobra.Command{
	Use:   "send-keys <pane_id> <key> [key]",
	Short: "Send key presses to a pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pane_sendKeysCmd).Standalone()

	paneCmd.AddCommand(pane_sendKeysCmd)

	carapace.Gen(pane_sendKeysCmd).PositionalCompletion(herdr.ActionPanes(herdr.PaneOpts{}))
}
