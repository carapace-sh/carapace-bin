package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var config_deleteContextCmd = &cobra.Command{
	Use:   "delete-context NAME",
	Short: "Delete the specified context from the kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteContextCmd).Standalone()

	configCmd.AddCommand(config_deleteContextCmd)

	carapace.Gen(config_deleteContextCmd).PositionalCompletion(
		kubectl.ActionContexts(),
	)
}
