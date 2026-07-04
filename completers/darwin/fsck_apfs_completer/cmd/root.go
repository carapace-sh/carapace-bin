package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fsck_apfs",
	Short: "APFS filesystem consistency check",
	Long:  "https://keith.github.io/xcode-manpages/fsck_apfs.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("D", "D", false, "Debug mode")
	rootCmd.Flags().BoolS("S", "S", false, "Verify only")
	rootCmd.Flags().BoolS("T", "T", false, "Test mode")
	rootCmd.Flags().BoolS("W", "W", false, "W flag")
	rootCmd.Flags().BoolS("d", "d", false, "Debug output")
	rootCmd.Flags().BoolS("l", "l", false, "Live filesystem")
	rootCmd.Flags().BoolS("n", "n", false, "Assume no to all questions")
	rootCmd.Flags().BoolS("o", "o", false, "Optimize")
	rootCmd.Flags().BoolS("q", "q", false, "Quiet mode")
	rootCmd.Flags().BoolS("s", "s", false, "Strict mode")
	rootCmd.Flags().BoolS("y", "y", false, "Assume yes to all questions")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
