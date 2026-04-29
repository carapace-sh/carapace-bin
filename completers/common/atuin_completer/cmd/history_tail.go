package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Stream history events from the daemon as they are received",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_tailCmd).Standalone()

	history_tailCmd.Flags().BoolP("help", "h", false, "Print help")
	historyCmd.AddCommand(history_tailCmd)
}
