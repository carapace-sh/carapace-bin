package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "nbtstat",
	Short: "display NetBIOS over TCP/IP (NetBT) protocol statistics",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/nbtstat",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("R", "R", false, "purge and reload the name cache")
	rootCmd.Flags().BoolP("S", "S", false, "display sessions table with destination IP addresses")
	rootCmd.Flags().BoolP("a", "a", false, "display remote computer name table")
	rootCmd.Flags().BoolP("n", "n", false, "display local computer name table")
	rootCmd.Flags().BoolP("r", "r", false, "display name resolution statistics")
	rootCmd.Flags().BoolP("s", "s", false, "display sessions table")
}
