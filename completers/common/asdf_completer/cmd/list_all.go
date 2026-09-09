package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/asdf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var list_allCmd = &cobra.Command{
	Use:   "all",
	Short: "List all versions of a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(list_allCmd).Standalone()

	listCmd.AddCommand(list_allCmd)

	carapace.Gen(list_allCmd).PositionalCompletion(
		action.ActionPlugins(),
	)
}
