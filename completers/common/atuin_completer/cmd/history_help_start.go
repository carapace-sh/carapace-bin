package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_help_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Begins a new command in the history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_help_startCmd).Standalone()

	history_helpCmd.AddCommand(history_help_startCmd)
}
