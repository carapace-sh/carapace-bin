package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var access_setCmd = &cobra.Command{
	Use:   "set",
	Short: "set access status or mfa",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(access_setCmd).Standalone()
	accessCmd.AddCommand(access_setCmd)

	carapace.Gen(access_setCmd).PositionalCompletion(
		carapace.ActionValues(
			"status=public",
			"status=private",
			"mfa=none",
			"mfa=publish",
			"mfa=automation",
			"2fa=none",
			"2fa=publish",
			"2fa=automation",
		),
		action.ActionPackages(access_setCmd),
	)
}
