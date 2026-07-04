package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "newfs_apfs",
	Short: "Create APFS filesystem",
	Long:  "https://keith.github.io/xcode-manpages/newfs_apfs.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().BoolS("All|-A", "All|-A", false, "")
	rootCmd.Flags().StringS("Blocksize|-b", "Blocksize|-b", "", "")
	rootCmd.Flags().BoolS("Container|-C", "Container|-C", false, "")
	rootCmd.Flags().BoolS("Empty", "Empty", false, "")
	rootCmd.Flags().BoolS("Interactive|-i", "Interactive|-i", false, "")
	rootCmd.Flags().StringS("Name|-v", "Name|-v", "", "")
	rootCmd.Flags().StringS("Options|-o", "Options|-o", "", "")
	rootCmd.Flags().StringS("Quota|-q", "Quota|-q", "", "")
	rootCmd.Flags().StringS("Reserve|-r", "Reserve|-r", "", "")
	rootCmd.Flags().StringS("Volumesize|-s", "Volumesize|-s", "", "")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
