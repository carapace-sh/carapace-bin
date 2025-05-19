package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gum",
	Short: "A tool for glamorous shell scripts",
	Long:  "https://github.com/charmbracelet/gum",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.PersistentFlags().BoolP("help", "h", false, "Show context-sensitive help.")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print the version number")
}
