package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dirs",
	Short: "Display directory stack",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Directory-Stack-Builtins.html#index-dirs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("c", "c", false, "clear the directory stack by deleting all of the elements")
	rootCmd.Flags().BoolS("l", "l", false, "do not print tilde-prefixed versions of directories relative to your home directory")
	rootCmd.Flags().BoolS("p", "p", false, "print the directory stack with one entry per line")
	rootCmd.Flags().BoolS("v", "v", false, "print the directory stack with one entry per line prefixed with its position in the stack")
}
