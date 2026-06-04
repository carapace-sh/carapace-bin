package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "help",
	Short: "Display information about builtin commands",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-help",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("d", "d", false, "output short description for each pattern match")
	rootCmd.Flags().BoolS("m", "m", false, "display usage in pseudo-manpage format")
	rootCmd.Flags().BoolS("p", "p", false, "list all builtins matching pattern")
	rootCmd.Flags().BoolS("s", "s", false, "output short description for each builtin")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		shell.ActionBuiltins().FilterArgs(),
	)
}
