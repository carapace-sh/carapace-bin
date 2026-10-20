package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var toolAlias_unsetCmd = &cobra.Command{
	Use:     "unset",
	Short:   "Remove a backend alias or a version alias from the global config",
	Aliases: []string{"rm", "remove", "delete", "del"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolAlias_unsetCmd).Standalone()

	toolAlias_unsetCmd.Flags().BoolP("help", "h", false, "Print help")
	toolAliasCmd.AddCommand(toolAlias_unsetCmd)

	carapace.Gen(toolAlias_unsetCmd).PositionalCompletion(
		mise.ActionTools(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return mise.ActionToolAliases(c.Args[0])
		}),
	)
}
