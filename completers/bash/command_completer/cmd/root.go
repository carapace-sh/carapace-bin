package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "command",
	Short: "Execute a simple command or display information about commands",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("V", "V", false, "print a more verbose description of each command")
	rootCmd.Flags().BoolS("p", "p", false, "use a default value for PATH that is guaranteed to find all of the standard utilities")
	rootCmd.Flags().BoolS("v", "v", false, "print a single word indicating the command or filename that invokes command")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		shell.ActionExecutables().FilterArgs(),
	)
}
