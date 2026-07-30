package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ping",
	Short: "verify IP-level connectivity to another TCP/IP computer",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/ping",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("a", "a", false, "resolve addresses to hostnames")
	rootCmd.Flags().BoolP("f", "f", false, "set Don't Fragment flag in packet")
	rootCmd.Flags().IntP("i", "i", 0, "set Time To Live")
	rootCmd.Flags().IntP("l", "l", 0, "set send buffer size")
	rootCmd.Flags().IntP("n", "n", 0, "number of echo requests to send")
	rootCmd.Flags().BoolP("t", "t", false, "ping the specified host until stopped")
	rootCmd.Flags().IntP("v", "v", 0, "set Type Of Service")
	rootCmd.Flags().IntP("w", "w", 0, "timeout in milliseconds to wait for each reply")
}
