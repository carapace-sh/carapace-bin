package cmd

import (
	"github.com/rsteube/carapace"
	git "github.com/rsteube/carapace-bin/completers/git_completer/cmd"
	"github.com/spf13/cobra"
)

var grepCmd = &cobra.Command{
	Use:                "grep",
	Short:              "Print lines matching a pattern",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(grepCmd).Standalone()

	rootCmd.AddCommand(grepCmd)

	carapace.Gen(grepCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			c.Args = append([]string{"grep"}, c.Args...)
			return git.ActionExecute().Chdir(grepCmd.Root().Flag("C").Value.String()).Invoke(c).ToA()
		}),
	)
}
