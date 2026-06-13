package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "whence",
	Short: "Identify command types",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-whence",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("c", "c", false, "list all commands matching")
	rootCmd.Flags().BoolS("f", "f", false, "use functions")
	rootCmd.Flags().BoolS("v", "v", false, "verbose output")
	rootCmd.Flags().BoolS("w", "w", false, "describe how word would be interpreted")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionExecutables(),
	)
}
