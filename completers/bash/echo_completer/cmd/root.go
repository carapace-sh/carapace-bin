package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "echo",
	Short: "Write arguments to the standard output",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-echo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("E", "E", false, "explicitly suppress interpretation of backslash escapes")
	rootCmd.Flags().BoolS("e", "e", false, "enable interpretation of backslash escapes")
	rootCmd.Flags().BoolS("n", "n", false, "do not append a newline")
}
