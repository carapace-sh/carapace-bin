package cmd

import (
	"github.com/rsteube/carapace"
	git "github.com/rsteube/carapace-bin/completers/git_completer/cmd"
	"github.com/spf13/cobra"
)

var reflogCmd = &cobra.Command{
	Use:                "reflog",
	Short:              "Manage reflog information",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(reflogCmd).Standalone()

	rootCmd.AddCommand(reflogCmd)

	carapace.Gen(reflogCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			c.Args = append([]string{"reflog"}, c.Args...)
			return git.ActionExecute().Chdir(logCmd.Root().Flag("C").Value.String()).Invoke(c).ToA()
		}),
	)
}
