package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var system_dialStdioCmd = &cobra.Command{
	Use:    "dial-stdio",
	Short:  "Proxy the stdio stream to the daemon connection. Should not be invoked manually.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_dialStdioCmd).Standalone()

	systemCmd.AddCommand(system_dialStdioCmd)
}
