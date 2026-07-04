package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "uname",
	Short: "display operating system name and information",
	Long:  "https://keith.github.io/xcode-manpages/uname.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("a", "a", false, "Show all")
	rootCmd.Flags().BoolS("m", "m", false, "Show machine type")
	rootCmd.Flags().BoolS("n", "n", false, "Show hostname")
	rootCmd.Flags().BoolS("o", "o", false, "Show operating system")
	rootCmd.Flags().BoolS("p", "p", false, "Show processor type")
	rootCmd.Flags().BoolS("r", "r", false, "Show OS release")
	rootCmd.Flags().BoolS("s", "s", false, "Show OS name")
	rootCmd.Flags().BoolS("v", "v", false, "Show OS version")
}
