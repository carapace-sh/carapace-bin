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

	rootCmd.Flags().BoolS("A", "A", false, "Add all existing volumes")
	rootCmd.Flags().BoolS("C", "C", false, "Create a new container")
	rootCmd.Flags().BoolS("E", "E", false, "Enable encryption")

	rootCmd.Flags().StringS("R", "R", "", "Volume role")
	rootCmd.Flags().StringS("b", "b", "", "Block size")
	rootCmd.Flags().StringS("i", "i", "", "Interactive")
	rootCmd.Flags().StringS("o", "o", "", "Options")
	rootCmd.Flags().StringS("q", "q", "", "Volume quota")
	rootCmd.Flags().StringS("r", "r", "", "Volume reserve")
	rootCmd.Flags().StringS("s", "s", "", "Volume size")
	rootCmd.Flags().StringS("v", "v", "", "Volume name")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
