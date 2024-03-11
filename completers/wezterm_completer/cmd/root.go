package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wezterm [OPTIONS] [COMMAND]",
	Short: "Wez's Terminal Emulator",
	Long:  "http://github.com/wez/wezterm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().String("config", "", "Override specific configuration values")
	rootCmd.Flags().String("config-file", "", "Specify the configuration file to use, overrides the normal")
	rootCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.Flags().BoolP("skip-config", "n", false, "Skip loading wezterm.lua")
	rootCmd.Flags().BoolP("version", "V", false, "Print version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		// TODO config: https://wezfurlong.org/wezterm/config/lua/config/index.html
		"config-file": carapace.ActionFiles(),
	})
}
