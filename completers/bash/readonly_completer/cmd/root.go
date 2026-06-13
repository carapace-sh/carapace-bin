package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "readonly",
	Short: "Mark shell variables as unchangeable",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-readonly",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("A", "A", false, "each name refers to an associative array variable")
	rootCmd.Flags().BoolS("a", "a", false, "each name refers to an indexed array variable")
	rootCmd.Flags().BoolS("f", "f", false, "each name refers to a shell function")
	rootCmd.Flags().BoolS("p", "p", false, "display a list of all readonly variables or functions")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return shell.ActionVariables()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
