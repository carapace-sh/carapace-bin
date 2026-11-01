package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vsdbutil",
	Short: "manipulates the volume status DB",
	Long:  "https://keith.github.io/xcode-manpages/vsdbutil.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("a", "a", false, "Adopt on-disk ownership on the specified path")
	rootCmd.Flags().BoolS("c", "c", false, "Check the status of the ownership usage on the specified path")
	rootCmd.Flags().BoolS("d", "d", false, "Disown on-disk ownership on the specified path")
	rootCmd.Flags().BoolS("h", "h", false, "Print help")
	rootCmd.Flags().BoolS("i", "i", false, "Initialize the ownership database")
	rootCmd.Flags().BoolS("x", "x", false, "Clear the entry associated with the specified path")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
