package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var alias_deleteCmd = &cobra.Command{
	Use:   "delete {<alias> | --all}",
	Short: "Delete set aliases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_deleteCmd).Standalone()

	alias_deleteCmd.Flags().Bool("all", false, "Delete all aliases")
	aliasCmd.AddCommand(alias_deleteCmd)

	carapace.Gen(alias_deleteCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if alias_deleteCmd.Flag("all").Changed {
				return carapace.ActionValues()
			}
			return action.ActionAliases()
		}),
	)
}
