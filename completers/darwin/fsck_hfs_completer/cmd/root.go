package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fsck_hfs",
	Short: "HFS+ filesystem consistency check",
	Long:  "https://keith.github.io/xcode-manpages/fsck_hfs.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolS("Debug|-d", "Debug|-d", false, "")
	rootCmd.Flags().BoolS("E", "E", false, "")
	rootCmd.Flags().BoolS("Extended|-x", "Extended|-x", false, "")
	rootCmd.Flags().BoolS("Force|-f", "Force|-f", false, "")
	rootCmd.Flags().BoolS("G|-g", "G|-g", false, "")
	rootCmd.Flags().BoolS("Live|-l", "Live|-l", false, "")
	rootCmd.Flags().BoolS("No|-n", "No|-n", false, "")
	rootCmd.Flags().BoolS("Preen|-p", "Preen|-p", false, "")
	rootCmd.Flags().BoolS("Quiet|-q", "Quiet|-q", false, "")
	rootCmd.Flags().BoolS("Rebuild|-r", "Rebuild|-r", false, "")
	rootCmd.Flags().BoolS("Snapshot|-S", "Snapshot|-S", false, "")
	rootCmd.Flags().BoolS("Yes|-y", "Yes|-y", false, "")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
