package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "wdutil",
	Short: "Wireless Diagnostics command line utility",
	Long:  "https://keith.github.io/xcode-manpages/wdutil.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"diagnose", "Run diagnostic tests",
			"info", "Display Wi-Fi information",
			"log", "Enable or disable logging",
			"dump", "Dump Wi-Fi log buffer to file",
		),
	)
}