package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var terminal_session_observeCmd = &cobra.Command{
	Use:   "observe",
	Short: "Observe a terminal stream",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(terminal_session_observeCmd).Standalone()

	terminal_session_observeCmd.Flags().String("cols", "", "")
	terminal_session_observeCmd.Flags().String("rows", "", "")
	terminal_sessionCmd.AddCommand(terminal_session_observeCmd)
}
