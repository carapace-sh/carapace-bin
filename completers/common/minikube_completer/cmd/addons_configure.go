package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var addons_configureCmd = &cobra.Command{
	Use:   "configure ADDON_NAME",
	Short: "Configures the addon w/ADDON_NAME within minikube (example: minikube addons configure registry-creds). For a list of available addons use: minikube addons list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addons_configureCmd).Standalone()

	addons_configureCmd.Flags().StringP("config-file", "f", "", "An optional configuration file to read addon specific configs from instead of being prompted each time.")
	addonsCmd.AddCommand(addons_configureCmd)

	carapace.Gen(addons_configureCmd).FlagCompletion(carapace.ActionMap{
		"config-file": carapace.ActionFiles(),
	})

	carapace.Gen(addons_configureCmd).PositionalCompletion(
		action.ActionAddons(),
	)
}
