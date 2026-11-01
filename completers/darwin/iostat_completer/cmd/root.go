package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iostat",
	Short: "report I/O statistics",
	Long:  "https://keith.github.io/xcode-manpages/iostat.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("C", "C", false, "Display CPU statistics")
	rootCmd.Flags().BoolS("I", "I", false, "Display total statistics for a given time period")
	rootCmd.Flags().BoolS("K", "K", false, "Display block count in kilobytes")
	rootCmd.Flags().BoolS("T", "T", false, "Display TTY statistics")
	rootCmd.Flags().BoolS("U", "U", false, "Display system load averages")
	rootCmd.Flags().StringS("c", "c", "", "Repeat the display count times")
	rootCmd.Flags().BoolS("d", "d", false, "Display only device statistics")
	rootCmd.Flags().StringS("n", "n", "", "Display up to devs number of devices")
	rootCmd.Flags().BoolS("o", "o", false, "Display old-style iostat device statistics")
	rootCmd.Flags().StringS("w", "w", "", "Wait interval")
}
