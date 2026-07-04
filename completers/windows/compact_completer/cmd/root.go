package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "compact",
	Short: "display or change the compression state of files or directories on NTFS partitions",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/compact",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("c", "c", false, "compress the specified files")
	rootCmd.Flags().BoolP("s", "s", false, "do not perform recursive operations")
	rootCmd.Flags().BoolP("u", "u", false, "uncompress the specified files")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
