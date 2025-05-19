package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var config_useContextCmd = &cobra.Command{
	Use:     "use-context CONTEXT_NAME",
	Short:   "Set the current-context in a kubeconfig file",
	Aliases: []string{"use"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_useContextCmd).Standalone()

	configCmd.AddCommand(config_useContextCmd)

	carapace.Gen(config_useContextCmd).PositionalCompletion(
		kubectl.ActionContexts(),
	)
}
