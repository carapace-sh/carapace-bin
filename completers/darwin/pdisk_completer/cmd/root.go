package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pdisk",
	Short: "Apple partition table editor",
	Long:  "https://keith.github.io/xcode-manpages/pdisk.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("abbr", "a", false, "Abbreviate the partition types shown in the partition list")
	rootCmd.Flags().BoolP("compute_size", "c", false, "Ignore the device size listed in the partition table")
	rootCmd.Flags().BoolP("debug", "d", false, "Turn on debugging")
	rootCmd.Flags().BoolP("fname", "f", false, "Show HFS volume names instead of partition name")
	rootCmd.Flags().BoolP("help", "h", false, "Print a short help message")
	rootCmd.Flags().BoolP("interactive", "i", false, "Go into interactive mode")
	rootCmd.Flags().BoolP("list", "l", false, "List partition tables for the specified devices")
	rootCmd.Flags().BoolP("logical", "L", false, "Show partition limits in logical blocks")
	rootCmd.Flags().BoolP("readonly", "r", false, "Prevent pdisk from writing to the device")
	rootCmd.Flags().BoolP("version", "v", false, "Print the version number")
}
