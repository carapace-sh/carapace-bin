package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gatsby_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Useful commands relating to Gatsby plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pluginCmd).Standalone()

	rootCmd.AddCommand(pluginCmd)

	carapace.Gen(pluginCmd).PositionalCompletion(
		carapace.ActionValues("docs", "ls"),
	)

	carapace.Gen(pluginCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionPlugins().Invoke(c).Filter(c.Args[1:]...).ToA() // TODO use filterargs with shift(1)
		}),
	)
}
