package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "type",
	Short: "Display information about command type",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-type",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("P", "P", false, "force a PATH search for each name")
	rootCmd.Flags().BoolS("a", "a", false, "display all locations containing an executable named name")
	rootCmd.Flags().BoolS("f", "f", false, "suppress shell function lookup")
	rootCmd.Flags().BoolS("p", "p", false, "return the name of the disk file that would be executed")
	rootCmd.Flags().BoolS("t", "t", false, "output a single word which is one of alias, keyword, function, builtin, file, or empty")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			// TODO shell.ActionExecutables currently bridges CarapaceBin for further args (good or bad?)
			return shell.ActionExecutables().Shift(len(c.Args)).FilterArgs()
		}),
	)
}
