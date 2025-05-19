package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/circleci_completer/cmd/action"
	"github.com/spf13/cobra"
)

var orb_removeFromCategoryCmd = &cobra.Command{
	Use:   "remove-from-category",
	Short: "Remove an orb from a category",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orb_removeFromCategoryCmd).Standalone()
	orbCmd.AddCommand(orb_removeFromCategoryCmd)

	carapace.Gen(orb_removeFromCategoryCmd).PositionalCompletion(
		action.ActionOrbs(),
		action.ActionOrbCategories(),
	)
}
