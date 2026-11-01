package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rex"
	"github.com/spf13/cobra"
)

var moveCmd = &cobra.Command{
	Use:   "move [BLOCK] DIRECTION ANCHOR",
	Short: "Move a block next to an anchor block",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moveCmd).Standalone()

	moveCmd.Flags().String("ratio", "", "give the moved block PCT percent of the split")

	common.AddWindowScopeFlags(moveCmd)

	rootCmd.AddCommand(moveCmd)

	carapace.Gen(moveCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch len(c.Args) {
			case 1:
				return rex.ActionDirections()
			default:
				return action.ActionBlocks(moveCmd)
			}
		}),
	)
}
