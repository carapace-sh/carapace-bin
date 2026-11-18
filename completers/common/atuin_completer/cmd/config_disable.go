package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable a feature, leaving its dependencies enabled",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_disableCmd).Standalone()

	config_disableCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.AddCommand(config_disableCmd)

	carapace.Gen(config_disableCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"daemon", "The background daemon and daemon-backed search",
			"output-capture", "Command output capture, via the daemon and the pty proxy",
			"threads", "Capture your AI coding agents' threads, to search and resume them",
		),
	)
}
