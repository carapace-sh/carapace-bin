package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset values and attributes of shell variables and functions",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-unset",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("f", "f", false, "treat each name as a shell function")
	rootCmd.Flags().BoolS("n", "n", false, "treat each name as a name reference and unset the variable itself")
	rootCmd.Flags().BoolS("v", "v", false, "treat each name as a shell variable")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flags().Lookup("f").Changed {
				return shell.ActionFunctions(true)
			}
			return shell.ActionVariables()
		}).FilterArgs(),
	)
}
