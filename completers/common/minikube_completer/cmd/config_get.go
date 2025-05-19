package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets the value of PROPERTY_NAME from the minikube config file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getCmd).Standalone()
	configCmd.AddCommand(config_getCmd)

	carapace.Gen(config_getCmd).PositionalCompletion(
		action.ActionConfigNames(),
	)
}
