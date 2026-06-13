package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "readonly",
	Short: "Mark variables as readonly",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-readonly",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("f", "f", false, "refer to shell functions")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		shell.ActionVariables(),
	)
}
