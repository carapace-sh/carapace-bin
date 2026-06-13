package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "alias",
	Short: "Define or display aliases",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-alias",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	// TODO missing flags and support for `+flag` notation
	rootCmd.Flags().BoolS("g", "g", false, "define a global alias")
	rootCmd.Flags().BoolS("s", "s", false, "define a suffix alias")

	carapace.Gen(rootCmd).PositionalCompletion(
		shell.ActionAliases(),
	)
}
