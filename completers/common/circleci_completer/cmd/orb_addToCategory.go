package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/circleci_completer/cmd/action"
	"github.com/spf13/cobra"
)

var orb_addToCategoryCmd = &cobra.Command{
	Use:   "add-to-category",
	Short: "Add an orb to a category",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orb_addToCategoryCmd).Standalone()
	orbCmd.AddCommand(orb_addToCategoryCmd)

	carapace.Gen(orb_addToCategoryCmd).PositionalCompletion(
		action.ActionOrbs(),
		action.ActionOrbCategories(),
	)
}
