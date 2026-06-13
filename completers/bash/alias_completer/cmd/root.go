package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "alias",
	Short: "Define or display aliases",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-alias",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("p", "p", false, "print all defined aliases in a reusable format")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return shell.ActionAliases()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
