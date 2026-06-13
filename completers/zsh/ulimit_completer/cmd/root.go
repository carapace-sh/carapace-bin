package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ulimit",
	Short: "Set or display resource limits",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-ulimit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("H", "H", false, "hard limit")
	rootCmd.Flags().BoolS("S", "S", false, "soft limit")
	rootCmd.Flags().BoolS("a", "a", false, "all limits")
}
