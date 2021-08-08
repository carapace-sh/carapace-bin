package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pub",
	Short: "Pub is a package manager for Dart",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	rootCmd.Flags().Bool("no-trace", false, "Do not print debugging information when an error occurs.")
	rootCmd.Flags().Bool("trace", false, "Print debugging information when an error occurs.")
	rootCmd.Flags().BoolP("verbose", "v", false, "Shortcut for \"--verbosity=all\".")
}
