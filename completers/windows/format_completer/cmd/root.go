package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "format",
	Short: "format a disk volume with a specified file system",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("q", "q", false, "quick format")
	rootCmd.Flags().BoolP("v", "v", false, "specify volume label")
}
