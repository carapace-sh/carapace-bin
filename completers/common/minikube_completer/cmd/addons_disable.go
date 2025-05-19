package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var addons_disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disables the addon w/ADDON_NAME within minikube (example: minikube addons disable dashboard). For a list of available addons use: minikube addons list ",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addons_disableCmd).Standalone()
	addonsCmd.AddCommand(addons_disableCmd)

	carapace.Gen(addons_disableCmd).PositionalCompletion(
		action.ActionAddons(),
	)
}
