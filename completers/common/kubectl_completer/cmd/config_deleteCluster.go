package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var config_deleteClusterCmd = &cobra.Command{
	Use:   "delete-cluster NAME",
	Short: "Delete the specified cluster from the kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_deleteClusterCmd).Standalone()

	configCmd.AddCommand(config_deleteClusterCmd)

	carapace.Gen(config_deleteClusterCmd).PositionalCompletion(
		kubectl.ActionClusters(),
	)
}
