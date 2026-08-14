package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_writeCharsCmd = &cobra.Command{
	Use:   "write-chars",
	Short: "Write characters to the terminal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_writeCharsCmd).Standalone()

	help_actionCmd.AddCommand(help_action_writeCharsCmd)
}
