package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/shell"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable and disable shell builtins",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Bash-Builtins.html#index-enable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("a", "a", false, "print a list of builtins showing whether or not each is enabled")
	rootCmd.Flags().BoolS("d", "d", false, "remove a builtin loaded with -f")
	rootCmd.Flags().StringS("f", "f", "", "load builtin name from shared object filename")
	rootCmd.Flags().BoolS("n", "n", false, "disable each name or display a list of disabled builtins")
	rootCmd.Flags().BoolS("p", "p", false, "print the list of builtins in a reusable format")
	rootCmd.Flags().BoolS("s", "s", false, "print only the names of Posix special builtins")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"f": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		shell.ActionBuiltins().FilterArgs(),
	)
}
