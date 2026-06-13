package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "history",
	Short: "Display or manipulate the command history",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("d", "d", false, "delete the specified history entry")
	rootCmd.Flags().BoolS("n", "n", false, "don't include event numbers")
	rootCmd.Flags().BoolS("p", "p", false, "perform history expansion")
	rootCmd.Flags().BoolS("r", "r", false, "read the history file")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionValues(),
	)
}
