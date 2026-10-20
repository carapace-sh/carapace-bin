package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var toolAlias_setCmd = &cobra.Command{
	Use:     "set",
	Short:   "Set a backend alias or a version alias in the global config",
	Aliases: []string{"add", "create"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolAlias_setCmd).Standalone()

	toolAlias_setCmd.Flags().BoolP("help", "h", false, "Print help")
	toolAliasCmd.AddCommand(toolAlias_setCmd)

	carapace.Gen(toolAlias_setCmd).PositionalCompletion(
		mise.ActionTools(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return mise.ActionToolAliases(c.Args[0])
		}),
	)
}
