package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "netstat",
	Short: "display active TCP connections, ports, and statistics",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/netstat",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("a", "a", false, "display all active TCP connections and ports")
	rootCmd.Flags().BoolP("b", "b", false, "display executable involved in each connection")
	rootCmd.Flags().BoolP("e", "e", false, "display Ethernet statistics")
	rootCmd.Flags().BoolP("n", "n", false, "display addresses and port numbers numerically")
	rootCmd.Flags().BoolP("o", "o", false, "display owning process ID for each connection")
	rootCmd.Flags().BoolP("r", "r", false, "display the routing table")
	rootCmd.Flags().BoolP("s", "s", false, "display statistics by protocol")
}
