package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/ufw_completer/cmd/actions"
	"github.com/spf13/cobra"
)

var app_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "show information on PROFILE",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(app_infoCmd)

	carapace.Gen(app_infoCmd).PositionalCompletion(
		actions.ActionUfwProfiles(),
	)

	appCmd.AddCommand(app_infoCmd)
}
