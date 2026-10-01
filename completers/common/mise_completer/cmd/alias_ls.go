package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var alias_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List tool version aliases",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_lsCmd).Standalone()

	aliasCmd.AddCommand(alias_lsCmd)

	carapace.Gen(alias_lsCmd).PositionalCompletion(
		action.ActionTools(),
	)
}
