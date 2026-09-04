package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/asdf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed versions of a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	rootCmd.AddCommand(listCmd)

	carapace.Gen(listCmd).PositionalCompletion(
		action.ActionPlugins(),
	)
}
