package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "umask",
	Short: "Display or set file creation mode mask",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-umask",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("S", "S", false, "symbolic mode")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues(
			"0", "1", "2", "3", "4", "5", "6", "7",
		).Usage("octal digit"),
	)
}
