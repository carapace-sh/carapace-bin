package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Display values currently set in the minikube config file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_viewCmd).Standalone()
	config_viewCmd.Flags().String("format", "- {{.ConfigKey}}: {{.ConfigValue}}", "Go template format string for the config view output.")
	configCmd.AddCommand(config_viewCmd)

	carapace.Gen(config_viewCmd).PositionalAnyCompletion(
		action.ActionConfigNames(),
	)
}
