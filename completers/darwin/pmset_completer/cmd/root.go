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

	rootCmd.Flags().BoolP("ac", "a", false, "Apply settings to AC power")
	rootCmd.Flags().BoolP("battery", "b", false, "Apply settings to battery power")
	rootCmd.Flags().BoolP("charger", "c", false, "Apply settings to charger power")
	rootCmd.Flags().BoolP("get", "g", false, "Get the current power management settings")
	rootCmd.Flags().BoolP("ups", "u", false, "Apply settings to UPS power")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"batt", "Show battery status",
			"chargers", "Show charger status",
			"custom", "Show custom settings",
			"live", "Show live settings",
			"assertions", "Show power assertions",
			"log", "Show power management log",
			"stats", "Show power management statistics",
			"systemstate", "Show system power state",
			"therm", "Show thermal warnings",
			"thermlog", "Show thermal log",
			"ac", "Show AC adapter information",
			"pslog", "Show power source log",
			"rawbatt", "Show raw battery data",
			"uuid", "Show the current UUID",
			"everything", "Show all available information",
		),
	)
}
