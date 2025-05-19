package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var alias_setCmd = &cobra.Command{
	Use:   "set <alias name> '<command>' [flags]",
	Short: "Set an alias for a longer command.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_setCmd).Standalone()

	alias_setCmd.Flags().BoolP("shell", "s", false, "Declare an alias to be passed through a shell interpreter.")
	aliasCmd.AddCommand(alias_setCmd)

	carapace.Gen(alias_setCmd).PositionalCompletion(
		action.ActionAliases(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if !alias_setCmd.Flag("shell").Changed {
				return bridge.ActionCarapaceBin("glab").Split()
			}
			return bridge.ActionCarapaceBin().SplitP()
		}),
	)
}
