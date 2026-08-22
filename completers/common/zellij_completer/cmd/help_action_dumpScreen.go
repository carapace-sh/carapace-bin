package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_dumpScreenCmd = &cobra.Command{
	Use:   "dump-screen",
	Short: "Dumps the viewport and optionally scrollback of a pane to a file or STDOUT",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_dumpScreenCmd).Standalone()

	help_actionCmd.AddCommand(help_action_dumpScreenCmd)
}
