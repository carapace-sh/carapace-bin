package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write bytes to the terminal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_writeCmd).Standalone()

	help_actionCmd.AddCommand(help_action_writeCmd)
}
