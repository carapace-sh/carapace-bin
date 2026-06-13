package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset variables and functions",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-unset",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("f", "f", false, "refer to shell functions")
	rootCmd.Flags().BoolS("v", "v", false, "refer to shell variables")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		shell.ActionVariables().FilterArgs(),
	)
}
