package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_dumpScreenCmd = &cobra.Command{
	Use:   "dump-screen",
	Short: "Dumps the viewport and optionally scrollback of a pane to a file or STDOUT",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_dumpScreenCmd).Standalone()

	action_helpCmd.AddCommand(action_help_dumpScreenCmd)
}
