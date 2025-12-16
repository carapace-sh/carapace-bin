package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/bluetoothctl"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bluetoothctl",
	Short: "Bluetooth Control Command Line Tool",
	Long:  "https://www.bluez.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("agent", "a", "", "Register agent handler: <capability>")
	rootCmd.Flags().BoolP("endpoints", "e", false, "Register media endpoints")
	rootCmd.Flags().BoolP("help", "h", false, "Display help")
	rootCmd.Flags().BoolP("monitor", "m", false, "Enable monitor output")
	rootCmd.Flags().StringP("timeout", "t", "", "Timeout in seconds for non-interactive mode")
	rootCmd.Flags().BoolP("version", "v", false, "Display version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"agent": bluetoothctl.ActionAgentCapabilities(),
	})
}
