package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_state_logs_profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Performance profile from a trace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_state_logs_profileCmd).Standalone()

	config_state_logs_profileCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	config_state_logsCmd.AddCommand(config_state_logs_profileCmd)

	carapace.Gen(config_state_logs_profileCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)

}
