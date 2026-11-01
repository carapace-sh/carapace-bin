package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "whois",
	Short: "Internet domain name and network number directory service",
	Long:  "https://keith.github.io/xcode-manpages/whois.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "Use the APNIC database")
	rootCmd.Flags().BoolS("I", "I", false, "Use the IANA database")
	rootCmd.Flags().BoolS("P", "P", false, "Use the PeeringDB database")
	rootCmd.Flags().BoolS("Q", "Q", false, "Do a quick lookup")
	rootCmd.Flags().BoolS("R", "R", false, "Do a recursive lookup")
	rootCmd.Flags().BoolS("S", "S", false, "Adjust simple queries for whois servers")
	rootCmd.Flags().BoolS("a", "a", false, "Use the ARIN database")
	rootCmd.Flags().BoolS("b", "b", false, "Use the Network Abuse Clearinghouse database")
	rootCmd.Flags().BoolS("f", "f", false, "Use the AfriNIC database")
	rootCmd.Flags().BoolS("g", "g", false, "Use the US non-military federal government database")
	rootCmd.Flags().BoolS("i", "i", false, "Use the InterNIC database")
	rootCmd.Flags().BoolS("k", "k", false, "Use the KRNIC database")
	rootCmd.Flags().BoolS("l", "l", false, "Use the LACNIC database")
	rootCmd.Flags().BoolS("m", "m", false, "Use the RADB database")
	rootCmd.Flags().BoolS("r", "r", false, "Use the RIPE database")

	rootCmd.Flags().StringS("h", "h", "", "Use the specified host as the whois server")
	rootCmd.Flags().StringS("p", "p", "", "Use the specified port for the whois server")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
