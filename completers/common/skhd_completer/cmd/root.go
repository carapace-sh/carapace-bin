package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

// TODO only relevant for darwin (https://github.com/carapace-sh/carapace-bin/issues/1326)
var rootCmd = &cobra.Command{
	Use:   "skhd",
	Short: "Simple hotkey daemon for macOS",
	Long:  "https://github.com/koekeishiya/skhd",
	Run:   func(*cobra.Command, []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("config", "c", "", "Specify location of config file")
	rootCmd.Flags().Bool("install-service", false, "Install launchd service file into ~/Library/LaunchAgents/com.koekeishiya.skhd.plist")
	rootCmd.Flags().StringP("key", "k", "", "Synthesize a keypress (same syntax as when defining a hotkey)")
	rootCmd.Flags().BoolP("no-hotload", "h", false, "Disable system for hotloading config file")
	rootCmd.Flags().BoolP("observe", "o", false, "Output keycode and modifiers of event. Ctrl+C to quit")
	rootCmd.Flags().BoolP("profile", "P", false, "Output profiling information")
	rootCmd.Flags().BoolP("reload", "r", false, "Signal a running instance of skhd to reload its config file")
	rootCmd.Flags().Bool("restart-service", false, "Restart skhd service")
	rootCmd.Flags().Bool("start-service", false, "Run skhd as a service through launchd")
	rootCmd.Flags().Bool("stop-service", false, "Stop skhd service from running")
	rootCmd.Flags().StringP("text", "t", "", "Synthesize a line of text")
	rootCmd.Flags().Bool("uninstall-service", false, "Remove launchd service file ~/Library/LaunchAgents/com.koekeishiya.skhd.plist")
	rootCmd.Flags().BoolP("verbose", "B", false, "Output debug information")
	rootCmd.Flags().BoolP("version", "v", false, "Print version number to stdout")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionFiles(),
	})
}
