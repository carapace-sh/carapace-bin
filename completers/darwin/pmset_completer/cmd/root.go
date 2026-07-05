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
	rootCmd.Flags().StringP("get", "g", "", "Get the current power management settings")
	rootCmd.Flags().BoolP("ups", "u", false, "Apply settings to UPS power")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"get": carapace.ActionValuesDescribed(
			"ac", "Show AC adapter information",
			"assertions", "Show power assertions",
			"batt", "Show battery status",
			"custom", "Show custom settings",
			"everything", "Show all available information",
			"live", "Show live settings",
			"log", "Show power management log",
			"pslog", "Show power source log",
			"rawbatt", "Show raw battery data",
			"stats", "Show power management statistics",
			"systemstate", "Show system power state",
			"therm", "Show thermal warnings",
			"thermlog", "Show thermal log",
			"uuid", "Show the current UUID",
		),
	})
}
