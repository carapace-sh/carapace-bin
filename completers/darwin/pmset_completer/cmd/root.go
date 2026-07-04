package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pmset",
	Short: "manipulate power management settings",
	Long:  "https://keith.github.io/xcode-manpages/pmset.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("batt", "b", false, "Show battery status")
	rootCmd.Flags().BoolP("chargers", "c", false, "Show charger status")
	rootCmd.Flags().BoolP("custom", "c", false, "Show custom settings")
	rootCmd.Flags().BoolP("help", "h", false, "Display usage information")
	rootCmd.Flags().BoolP("live", "l", false, "Show live settings")
	rootCmd.Flags().BoolP("touch", "t", false, "Touch the PM settings")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"-g", "Get the current power management settings",
			"-g batt", "Show battery status",
			"-g chargers", "Show charger status",
			"-g custom", "Show custom settings",
			"-g live", "Show live settings",
			"-g assertions", "Show power assertions",
			"-g log", "Show power management log",
			"-g stats", "Show power management statistics",
			"-g systemstate", "Show system power state",
			"-g therm", "Show thermal warnings",
			"-g thermlog", "Show thermal log",
			"-g ac", "Show AC adapter information",
			"-g pslog", "Show power source log",
			"-g rawbatt", "Show raw battery data",
			"-g uuid", "Show the current UUID",
			"-g everything", "Show all available information",
		),
	)
}
