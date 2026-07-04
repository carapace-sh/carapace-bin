package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "icacls",
	Short: "display or modify discretionary access control lists (DACLs)",
	Long:  "https://learn.microsoft.com/en-us/windows-server/administration/windows-commands/icacls",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("q", "q", false, "suppress success messages")
	rootCmd.Flags().BoolP("c", "c", false, "continue on access denied errors")
	rootCmd.Flags().BoolP("t", "t", false, "operate recursively on files in directory and subdirectories")
	rootCmd.Flags().BoolP("l", "l", false, "perform operation on a symbolic link vs. its target")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
