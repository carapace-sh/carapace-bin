package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_help_tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Stream history events from the daemon as they are received",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_help_tailCmd).Standalone()

	history_helpCmd.AddCommand(history_help_tailCmd)
}
