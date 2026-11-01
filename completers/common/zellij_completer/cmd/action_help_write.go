package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write bytes to the terminal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_writeCmd).Standalone()

	action_helpCmd.AddCommand(action_help_writeCmd)
}
