package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/bluetoothctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bluetoothctl",
	Short: "Bluetooth Control Command Line Tool",
	Long:  "https://www.bluez.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("agent", "a", "", "Register agent handler: <capability>")
	rootCmd.Flags().BoolP("endpoints", "e", false, "Register media endpoints")
	rootCmd.Flags().BoolP("monitor", "m", false, "Enable monitor output")
	rootCmd.Flags().BoolP("version", "v", false, "Display version")
	rootCmd.Flags().BoolP("help", "h", false, "Display help")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"agent": action.ActionAgents(),
	})
}
