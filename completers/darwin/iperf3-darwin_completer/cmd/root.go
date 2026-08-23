package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iperf3-darwin",
	Short: "perform network throughput tests",
	Long:  "https://man.freebsd.org/cgi/man.cgi?iperf3-darwin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("c", "c", false, "Run in client mode")
	rootCmd.Flags().BoolS("h", "h", false, "Help")
	rootCmd.Flags().BoolS("s", "s", false, "Run in server mode")
	rootCmd.Flags().BoolS("v", "v", false, "Version")
}
