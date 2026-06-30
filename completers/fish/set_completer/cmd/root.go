package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "set",
	Short: "Display and change shell variables",
	Long:  "https://fishshell.com/docs/current/cmds/set.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("L", "L", false, "don't abbreviate")
	rootCmd.Flags().BoolS("S", "S", false, "show variable info")
	rootCmd.Flags().BoolS("U", "U", false, "universal variable")
	rootCmd.Flags().BoolS("a", "a", false, "append values")
	rootCmd.Flags().Bool("append", false, "append values")
	rootCmd.Flags().BoolS("e", "e", false, "erase variables")
	rootCmd.Flags().Bool("erase", false, "erase variables")
	rootCmd.Flags().Bool("export", false, "export")
	rootCmd.Flags().BoolS("f", "f", false, "function scope")
	rootCmd.Flags().Bool("function", false, "function scope")
	rootCmd.Flags().BoolS("g", "g", false, "global scope")
	rootCmd.Flags().Bool("global", false, "global scope")
	rootCmd.Flags().BoolS("h", "h", false, "display help")
	rootCmd.Flags().BoolS("l", "l", false, "local scope")
	rootCmd.Flags().Bool("local", false, "local scope")
	rootCmd.Flags().Bool("long", false, "don't abbreviate")
	rootCmd.Flags().BoolS("n", "n", false, "list names")
	rootCmd.Flags().Bool("names", false, "list names")
	rootCmd.Flags().Bool("no-event", false, "no event")
	rootCmd.Flags().BoolS("p", "p", false, "prepend values")
	rootCmd.Flags().Bool("path", false, "path variable")
	rootCmd.Flags().Bool("prepend", false, "prepend values")
	rootCmd.Flags().BoolS("q", "q", false, "query variables")
	rootCmd.Flags().Bool("query", false, "query variables")
	rootCmd.Flags().Bool("show", false, "show variable info")
	rootCmd.Flags().BoolS("u", "u", false, "do not export")
	rootCmd.Flags().Bool("unexport", false, "do not export")
	rootCmd.Flags().Bool("universal", false, "universal variable")
	rootCmd.Flags().Bool("unpath", false, "unpath variable")
	rootCmd.Flags().BoolS("x", "x", false, "export")

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
