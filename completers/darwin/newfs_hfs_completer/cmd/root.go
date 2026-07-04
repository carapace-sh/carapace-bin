package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "newfs_hfs",
	Short: "Create HFS+ filesystem",
	Long:  "https://keith.github.io/xcode-manpages/newfs_hfs.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().StringS("Blocksize|-b", "Blocksize|-b", "", "")
	rootCmd.Flags().BoolS("CaseSensitive|-s", "CaseSensitive|-s", false, "")
	rootCmd.Flags().StringS("Clump|-c", "Clump|-c", "", "")
	rootCmd.Flags().StringS("Cnid|-i", "Cnid|-i", "", "")
	rootCmd.Flags().StringS("Gid|-G", "Gid|-G", "", "")
	rootCmd.Flags().StringS("Journaldev|-D", "Journaldev|-D", "", "")
	rootCmd.Flags().StringS("Journal|-J", "Journal|-J", "", "")
	rootCmd.Flags().StringS("Mask|-M", "Mask|-M", "", "")
	rootCmd.Flags().StringS("Name|-v", "Name|-v", "", "")
	rootCmd.Flags().BoolS("No", "No", false, "")
	rootCmd.Flags().StringS("Nodesize|-n", "Nodesize|-n", "", "")
	rootCmd.Flags().BoolS("Preserve|-P", "Preserve|-P", false, "")
	rootCmd.Flags().StringS("Uid|-U", "Uid|-U", "", "")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
