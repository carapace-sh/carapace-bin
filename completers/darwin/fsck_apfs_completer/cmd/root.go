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
	rootCmd.Flags().BoolS("Debug|-d", "Debug|-d", false, "")
	rootCmd.Flags().BoolS("Live|-l", "Live|-l", false, "")
	rootCmd.Flags().BoolS("No|-n", "No|-n", false, "")
	rootCmd.Flags().BoolS("Optimize|-o", "Optimize|-o", false, "")
	rootCmd.Flags().BoolS("Quiet|-q", "Quiet|-q", false, "")
	rootCmd.Flags().BoolS("Revert|-D", "Revert|-D", false, "")
	rootCmd.Flags().BoolS("Strict|-s", "Strict|-s", false, "")
	rootCmd.Flags().BoolS("T", "T", false, "")
	rootCmd.Flags().BoolS("Verify|-S", "Verify|-S", false, "")
	rootCmd.Flags().BoolS("W", "W", false, "")
	rootCmd.Flags().BoolS("Yes|-y", "Yes|-y", false, "")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
