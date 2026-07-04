package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pathping",
	Short: "provide information about network latency and packet loss",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/pathping",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("d", "d", false, "do not resolve addresses to hostnames")
	rootCmd.Flags().BoolP("h", "h", false, "maximum number of hops")
	rootCmd.Flags().BoolP("n", "n", false, "do not resolve addresses to hostnames")
	rootCmd.Flags().BoolP("q", "q", false, "number of queries per hop")
	rootCmd.Flags().BoolP("w", "w", false, "timeout for each reply")
}
