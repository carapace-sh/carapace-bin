package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "typeperf",
	Short: "write performance data to the command window or a file",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/typeperf",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
}
