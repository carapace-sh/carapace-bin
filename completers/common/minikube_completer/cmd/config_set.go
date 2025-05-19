package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets an individual value in a minikube config file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()
	configCmd.AddCommand(config_setCmd)

	carapace.Gen(config_setCmd).PositionalCompletion(
		action.ActionConfigNames(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionConfigDefaults(c.Args[0])
		}),
	)
}
