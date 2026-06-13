package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dirs",
	Short: "Display the directory stack",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-dirs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("c", "c", false, "clear the directory stack")
	rootCmd.Flags().BoolS("l", "l", false, "produce a longer listing")
	rootCmd.Flags().BoolS("p", "p", false, "print one entry per line")
	rootCmd.Flags().BoolS("v", "v", false, "print one entry per line with position index")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionValues(),
	)
}
