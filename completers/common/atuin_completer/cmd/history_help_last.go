package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_help_lastCmd = &cobra.Command{
	Use:   "last",
	Short: "Get the last command ran",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_help_lastCmd).Standalone()

	history_helpCmd.AddCommand(history_help_lastCmd)
}
