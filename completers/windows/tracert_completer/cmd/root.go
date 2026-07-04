package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tracert",
	Short: "determine the path taken to a destination",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/tracert",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("d", "d", false, "do not resolve IP addresses to hostnames")
	rootCmd.Flags().BoolP("h", "h", false, "maximum number of hops to search")
	rootCmd.Flags().BoolP("w", "w", false, "timeout in milliseconds for each reply")
}
