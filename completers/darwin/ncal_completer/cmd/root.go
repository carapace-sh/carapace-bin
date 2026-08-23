package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ncal",
	Short: "displays a calendar and the date of Easter",
	Long:  "https://man.freebsd.org/cgi/man.cgi?ncal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("3", "3", false, "Display previous, current and next month")
	rootCmd.Flags().StringS("A", "A", "", "Number of months after the current month")
	rootCmd.Flags().StringS("B", "B", "", "Number of months before the current month")
	rootCmd.Flags().BoolS("C", "C", false, "Switch to cal mode")
	rootCmd.Flags().StringS("H", "H", "", "Use yyyy-mm-dd for highlighting")
	rootCmd.Flags().BoolS("J", "J", false, "Display Julian Calendar")
	rootCmd.Flags().BoolS("N", "N", false, "Switch to ncal mode")
	rootCmd.Flags().StringS("d", "d", "", "Use yyyy-mm as the current date")
	rootCmd.Flags().BoolS("e", "e", false, "Display date of Easter")
	rootCmd.Flags().BoolS("h", "h", false, "Turns off highlighting of today")
	rootCmd.Flags().BoolS("j", "j", false, "Display Julian days")
	rootCmd.Flags().StringS("m", "m", "", "Display the specified month")
	rootCmd.Flags().BoolS("o", "o", false, "Display date of Orthodox Easter")
	rootCmd.Flags().BoolS("p", "p", false, "Print country codes")
	rootCmd.Flags().StringS("s", "s", "", "Assume switch date for country_code")
	rootCmd.Flags().BoolS("w", "w", false, "Print the number of the week")
	rootCmd.Flags().BoolS("y", "y", false, "Display a calendar for the specified year")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"m": carapace.ActionValues("1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"),
	})
}
