package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable a feature, along with everything it depends on",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_enableCmd).Standalone()

	config_enableCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	configCmd.AddCommand(config_enableCmd)

	carapace.Gen(config_enableCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"daemon", "The background daemon and daemon-backed search",
			"output-capture", "Command output capture, via the daemon and the pty proxy",
			"threads", "Capture your AI coding agents' threads, to search and resume them",
		),
	)
}
