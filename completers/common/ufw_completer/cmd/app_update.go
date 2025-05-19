package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/ufw_completer/cmd/actions"
	"github.com/spf13/cobra"
)

var app_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update PROFILE",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(app_updateCmd).Standalone()

	app_updateCmd.Flags().String("add-new", "", "")
	appCmd.AddCommand(app_updateCmd)

	carapace.Gen(app_updateCmd).PositionalCompletion(
		actions.ActionUfwProfiles(),
	)
}
