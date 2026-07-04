package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "portqry",
	Short: "port query tool for diagnosing TCP/IP connectivity",
	Long:  "https://learn.microsoft.com/en-us/troubleshoot/windows-server/networking/portqry-command-line-port-scanner",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("n", "n", false, "skip name resolution")
	rootCmd.Flags().BoolP("o", "o", false, "specify port(s) to query")
	rootCmd.Flags().BoolP("p", "p", false, "specify protocol (TCP or UDP)")
	rootCmd.Flags().BoolP("r", "r", false, "query range of ports")
}
