package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hidutil",
	Short: "HID device utility",
	Long:  "https://developer.apple.com/library/archive/documentation/DeviceDrivers/Conceptual/HID/new_api_10.5_intro/chapter_2_section_5.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"dump", "Dump HID Event System state",
			"help", "Print help message",
			"list", "List HID Event System services and devices",
			"property", "Read/write HID Event System property",
		),
	)
}
