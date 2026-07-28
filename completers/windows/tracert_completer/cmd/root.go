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
	rootCmd.Flags().IntP("h", "h", 0, "maximum number of hops to search")
	rootCmd.Flags().IntP("w", "w", 0, "timeout in milliseconds for each reply")
}
