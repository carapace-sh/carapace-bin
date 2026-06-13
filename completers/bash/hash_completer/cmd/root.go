package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hash",
	Short: "Remember or display program locations",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-hash",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("d", "d", false, "forget the remembered location of each name")
	rootCmd.Flags().BoolS("l", "l", false, "display in a format that may be reused as input")
	rootCmd.Flags().StringS("p", "p", "", "use pathname as the full path of name")
	rootCmd.Flags().BoolS("r", "r", false, "forget all remembered locations")
	rootCmd.Flags().StringS("t", "t", "", "print the remembered location of name")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"p": carapace.ActionFiles(),
		"t": shell.ActionExecutables(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		shell.ActionExecutables().FilterArgs(),
	)
}
