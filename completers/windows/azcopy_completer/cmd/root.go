package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "azcopy",
	Short: "command-line utility for copying data to/from Azure Storage",
	Long:  "https://learn.microsoft.com/en-us/azure/storage/common/storage-use-azcopy-v10",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
}
