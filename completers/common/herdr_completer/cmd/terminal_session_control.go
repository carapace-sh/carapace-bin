package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var terminal_session_controlCmd = &cobra.Command{
	Use:   "control",
	Short: "Control a terminal stream",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(terminal_session_controlCmd).Standalone()

	terminal_session_controlCmd.Flags().String("cols", "", "")
	terminal_session_controlCmd.Flags().String("rows", "", "")
	terminal_session_controlCmd.Flags().Bool("takeover", false, "")
	terminal_sessionCmd.AddCommand(terminal_session_controlCmd)
}
