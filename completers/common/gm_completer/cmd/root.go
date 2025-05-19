package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gm",
	Short: "command-line utility to create, edit, compare, convert, or display images",
	Long:  "http://www.GraphicsMagick.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("help", "help", false, "program options")
	rootCmd.Flags().BoolS("version", "version", false, "print version information")
}
