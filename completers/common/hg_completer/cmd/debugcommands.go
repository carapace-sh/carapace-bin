package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugcommandsCmd = &cobra.Command{
	Use:    "debugcommands",
	Short:  "list all available commands and options",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugcommandsCmd).Standalone()

	rootCmd.AddCommand(debugcommandsCmd)
}
