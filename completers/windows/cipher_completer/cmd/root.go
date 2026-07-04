package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cipher",
	Short: "display or alter the encryption of directories and files on NTFS volumes",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/cipher",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("d", "d", false, "decrypt the specified files")
	rootCmd.Flags().BoolP("e", "e", false, "encrypt the specified files")
	rootCmd.Flags().BoolP("s", "s", false, "perform the operation on all subdirectories")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
