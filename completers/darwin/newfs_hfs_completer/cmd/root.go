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

	rootCmd.Flags().BoolS("N", "N", false, "Do not create filesystem (show parameters)")
	rootCmd.Flags().BoolS("P", "P", false, "Preserve filesystem")
	rootCmd.Flags().BoolS("s", "s", false, "Case-sensitive filesystem")

	rootCmd.Flags().StringS("D", "D", "", "Journal device")
	rootCmd.Flags().StringS("G", "G", "", "Group ID")
	rootCmd.Flags().StringS("J", "J", "", "Journal size")
	rootCmd.Flags().StringS("M", "M", "", "Access mask")
	rootCmd.Flags().StringS("U", "U", "", "User ID")
	rootCmd.Flags().StringS("b", "b", "", "Block size")
	rootCmd.Flags().StringS("c", "c", "", "Clump size list")
	rootCmd.Flags().StringS("i", "i", "", "First CNID")
	rootCmd.Flags().StringS("n", "n", "", "Node size list")
	rootCmd.Flags().StringS("v", "v", "", "Volume name")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
