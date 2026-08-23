package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sntp",
	Short: "simple Network Time Protocol client program",
	Long:  "https://man.freebsd.org/cgi/man.cgi?sntp",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("S", "S", false, "Set the system clock")
	rootCmd.Flags().BoolS("d", "d", false, "Enable debug logging")
	rootCmd.Flags().StringS("g", "g", "", "Gap between requests in milliseconds")
	rootCmd.Flags().StringS("n", "n", "", "Number of DNS records to use")
	rootCmd.Flags().BoolS("r", "r", false, "Bind the NTP reserved port")
	rootCmd.Flags().BoolS("s", "s", false, "Slew the system clock")
	rootCmd.Flags().StringS("t", "t", "", "Maximum wait time in seconds")
	rootCmd.Flags().StringS("z", "z", "", "Path to dump header state")
}
