package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "copy",
	Short: "copy one or more files to another location",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("a", "a", false, "indicate an ASCII text file")
	rootCmd.Flags().BoolP("b", "b", false, "indicate a binary file")
	rootCmd.Flags().BoolP("d", "d", false, "allow the destination file to be created decrypted")
	rootCmd.Flags().BoolP("v", "v", false, "verify that files are written correctly")
	rootCmd.Flags().BoolP("y", "y", false, "suppress prompting to confirm overwriting")
	rootCmd.Flags().BoolP("z", "z", false, "copy files in restartable mode")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(),
	)
}
