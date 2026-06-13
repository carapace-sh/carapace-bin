package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "export",
	Short: "Set export attribute for shell variables",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-export",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("f", "f", false, "refer to shell functions")
	rootCmd.Flags().BoolS("n", "n", false, "remove the export property from each name")
	rootCmd.Flags().BoolS("p", "p", false, "display a list of all exported variables or functions")

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
