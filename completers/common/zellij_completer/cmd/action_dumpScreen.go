package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zellij"
	"github.com/spf13/cobra"
)

var action_dumpScreenCmd = &cobra.Command{
	Use:   "dump-screen",
	Short: "Dumps the viewport and optionally scrollback of a pane to a file or STDOUT",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_dumpScreenCmd).Standalone()

	action_dumpScreenCmd.Flags().BoolP("ansi", "a", false, "Preserve ANSI styling in the dump output")
	action_dumpScreenCmd.Flags().BoolP("full", "f", false, "Dump the pane with full scrollback")
	action_dumpScreenCmd.Flags().BoolP("help", "h", false, "Print help")
	action_dumpScreenCmd.Flags().StringP("pane-id", "p", "", "The pane_id of the pane, eg. terminal_1, plugin_2 or 3 (equivalent to terminal_3). If not specified, dumps the focused pane")
	action_dumpScreenCmd.Flags().String("path", "", "File path to dump the pane content to. If omitted, prints to STDOUT")
	actionCmd.AddCommand(action_dumpScreenCmd)

	carapace.Gen(action_dumpScreenCmd).FlagCompletion(carapace.ActionMap{
		"pane-id": zellij.ActionSelectablePanes(),
		"path":    carapace.ActionFiles(),
	})
}
