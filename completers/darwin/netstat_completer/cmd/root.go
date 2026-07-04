package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "netstat",
	Short: "Network statistics",
	Long:  "https://keith.github.io/xcode-manpages/netstat.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "Show addresses of control blocks associated with sockets")
	rootCmd.Flags().BoolS("L", "L", false, "Show listen queue lengths")
	rootCmd.Flags().BoolS("S", "S", false, "Show network interface send queue")
	rootCmd.Flags().BoolS("W", "W", false, "Wide output")
	rootCmd.Flags().BoolS("a", "a", false, "Show all sockets including server processes")
	rootCmd.Flags().BoolS("b", "b", false, "Show bytes in/out")
	rootCmd.Flags().BoolS("d", "d", false, "Show dropped packets")
	rootCmd.Flags().BoolS("g", "g", false, "Show multicast group memberships")
	rootCmd.Flags().BoolS("i", "i", false, "Show network interface statistics")
	rootCmd.Flags().BoolS("l", "l", false, "Show listening sockets")
	rootCmd.Flags().BoolS("m", "m", false, "Show memory statistics")
	rootCmd.Flags().BoolS("n", "n", false, "Show network addresses as numbers")
	rootCmd.Flags().BoolS("q", "q", false, "Show interface send queue")
	rootCmd.Flags().BoolS("r", "r", false, "Show routing tables")
	rootCmd.Flags().BoolS("s", "s", false, "Show network statistics for each protocol")
	rootCmd.Flags().BoolS("t", "t", false, "Show TCP connections")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose mode")

	rootCmd.Flags().StringS("I", "I", "", "Show interface statistics")
	rootCmd.Flags().StringS("R", "R", "", "Show route-references")
	rootCmd.Flags().StringS("c", "c", "", "Show queue length")
	rootCmd.Flags().StringS("f", "f", "", "Address family")
	rootCmd.Flags().StringS("p", "p", "", "Protocol")
	rootCmd.Flags().StringS("w", "w", "", "Wait time between updates")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionValues())
}
