package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_state_logs_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clear background operation logs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_logs_clearCmd).Standalone()

	config_state_logs_clearCmd.Flags().BoolP("help", "h", false, "Print help")
	config_state_logsCmd.AddCommand(config_state_logs_clearCmd)
}
