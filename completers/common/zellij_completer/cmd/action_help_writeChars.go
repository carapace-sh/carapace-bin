package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_writeCharsCmd = &cobra.Command{
	Use:   "write-chars",
	Short: "Write characters to the terminal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_writeCharsCmd).Standalone()

	action_helpCmd.AddCommand(action_help_writeCharsCmd)
}
