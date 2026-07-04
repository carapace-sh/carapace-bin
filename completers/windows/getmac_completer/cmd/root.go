package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "getmac",
	Short: "display the MAC address and list of network protocols",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/getmac",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("fo", "fo", false, "specify output format (table, list, csv)")
	rootCmd.Flags().BoolP("nh", "nh", false, "no column headers in output")
	rootCmd.Flags().BoolP("v", "v", false, "display verbose information")
}
