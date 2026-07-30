package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fsutil",
	Short: "file system utility for performing tasks related to FAT and NTFS file systems",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/fsutil",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
}
