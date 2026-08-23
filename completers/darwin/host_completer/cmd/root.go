package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "host",
	Short: "DNS lookup utility",
	Long:  "https://keith.github.io/xcode-manpages/host.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("4", "4", false, "Use IPv4 only for query transport")
	rootCmd.Flags().BoolS("6", "6", false, "Use IPv6 only for query transport")
	rootCmd.Flags().BoolS("C", "C", false, "Check consistency")
	rootCmd.Flags().StringS("N", "N", "", "Number of dots")
	rootCmd.Flags().StringS("R", "R", "", "Number of retries")
	rootCmd.Flags().BoolS("T", "T", false, "TCP mode")
	rootCmd.Flags().StringS("W", "W", "", "Wait time")
	rootCmd.Flags().BoolS("a", "a", false, "All")
	rootCmd.Flags().StringS("c", "c", "", "Query class")
	rootCmd.Flags().BoolS("d", "d", false, "Print debugging traces")
	rootCmd.Flags().BoolS("i", "i", false, "Obsolete")
	rootCmd.Flags().BoolS("l", "l", false, "List zone")
	rootCmd.Flags().StringS("m", "m", "", "Memory debugging flag")
	rootCmd.Flags().BoolS("n", "n", false, "Ignore reverse lookup")
	rootCmd.Flags().BoolS("r", "r", false, "Disable recursion")
	rootCmd.Flags().BoolS("s", "s", false, "Do not send queries to server")
	rootCmd.Flags().StringS("t", "t", "", "Query type")
	rootCmd.Flags().BoolS("v", "v", false, "Verbose")
	rootCmd.Flags().BoolS("w", "w", false, "Wait forever")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"c": carapace.ActionValues("IN", "HS", "CH"),
		"t": carapace.ActionValues("A", "AAAA", "CNAME", "MX", "NS", "PTR", "SOA", "SRV", "TXT", "ANY"),
	})
}
