package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hash",
	Short: "View or modify the command hash table",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-hash",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("L", "L", false, "print each hash table entry in the form of a call to hash")
	rootCmd.Flags().BoolS("d", "d", false, "use the named directory hash table instead of the command hash table")
	rootCmd.Flags().BoolS("f", "f", false, "cause the selected hash table to be fully rebuilt immediately")
	rootCmd.Flags().BoolS("m", "m", false, "take arguments as patterns and print matching hash table elements")
	rootCmd.Flags().BoolS("r", "r", false, "cause the selected hash table to be emptied")
	rootCmd.Flags().BoolS("v", "v", false, "list hash table entries as they are added by explicit specification")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionExecutables()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
