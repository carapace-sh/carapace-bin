package cmd

import (
	"github.com/rsteube/carapace"
	git "github.com/rsteube/carapace-bin/completers/git_completer/cmd"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:                "log",
	Short:              "Show commit logs",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(logCmd).Standalone()

	rootCmd.AddCommand(logCmd)

	carapace.Gen(logCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			c.Args = append([]string{"log"}, c.Args...)
			return git.ActionExecute().Chdir(logCmd.Root().Flag("C").Value.String()).Invoke(c).ToA()
		}),
	)
}
