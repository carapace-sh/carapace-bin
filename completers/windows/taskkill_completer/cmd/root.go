package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "taskkill",
	Short: "end one or more tasks or processes",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/taskkill",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("f", "f", false, "forcefully terminate the process")
	rootCmd.Flags().BoolP("t", "t", false, "terminate the process and any child processes")
}
