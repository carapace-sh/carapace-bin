package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "alias",
	Short: "Create a function",
	Long:  "https://fishshell.com/docs/current/cmds/alias.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().BoolS("s", "s", false, "save to fish configuration directory")
	rootCmd.Flags().Bool("save", false, "save to fish configuration directory")

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
