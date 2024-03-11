package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var config_renameContextCmd = &cobra.Command{
	Use:   "rename-context CONTEXT_NAME NEW_NAME",
	Short: "Rename a context from the kubeconfig file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_renameContextCmd).Standalone()

	configCmd.AddCommand(config_renameContextCmd)

	carapace.Gen(config_renameContextCmd).PositionalCompletion(
		kubectl.ActionContexts(),
	)
}
