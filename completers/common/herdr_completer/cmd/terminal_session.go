package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var terminal_sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Work with terminal sessions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(terminal_sessionCmd).Standalone()

	terminalCmd.AddCommand(terminal_sessionCmd)
}
