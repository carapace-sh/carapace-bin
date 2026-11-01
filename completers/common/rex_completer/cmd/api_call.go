package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rex"
	"github.com/spf13/cobra"
)

var apiCallCmd = &cobra.Command{
	Use:   "call METHOD [ARGS]",
	Short: "Call a server API method",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apiCallCmd).Standalone()

	apiCallCmd.Flags().StringP("session", "s", "", "scope to session SESSION")

	apiCmd.AddCommand(apiCallCmd)

	carapace.Gen(apiCallCmd).FlagCompletion(carapace.ActionMap{
		"session": action.ActionSessions(apiCallCmd),
	})

	carapace.Gen(apiCallCmd).PositionalCompletion(
		rex.ActionApiMethods(),
		carapace.Batch(
			carapace.ActionFiles().Prefix("@"),
			carapace.ActionValues("-"),
		).ToA(),
	)
}
