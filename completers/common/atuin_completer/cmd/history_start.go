package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_startCmd = &cobra.Command{
	Use:   "start",
	Short: "Begins a new command in the history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_startCmd).Standalone()

	history_startCmd.Flags().BoolP("help", "h", false, "Print help")
	historyCmd.AddCommand(history_startCmd)
}
