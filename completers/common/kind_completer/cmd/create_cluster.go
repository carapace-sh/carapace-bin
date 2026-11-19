package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/kind_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var create_clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Creates a local Kubernetes cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_clusterCmd).Standalone()

	create_clusterCmd.Flags().String("config", "", "path to a kind config file")
	create_clusterCmd.Flags().String("image", "", "node docker image to use for booting the cluster")
	create_clusterCmd.Flags().String("kubeconfig", "", "sets kubeconfig path instead of $KUBECONFIG or $HOME/.kube/config")
	create_clusterCmd.Flags().StringP("name", "n", "", "cluster name, overrides KIND_CLUSTER_NAME, config (default kind)")
	create_clusterCmd.Flags().Bool("retain", false, "retain nodes for debugging when cluster creation fails")
	create_clusterCmd.Flags().String("wait", "", "The length of time to wait for control plane node to be ready (default 0s). Non-zero values should contain a corresponding time unit (e.g. 1s, 2m, 3h).")

	carapace.Gen(create_clusterCmd).FlagCompletion(carapace.ActionMap{
		"config":     carapace.ActionFiles(".yaml", ".yml"),
		"image":      docker.ActionRepositoryTags(),
		"kubeconfig": carapace.ActionFiles(),
		"name":       action.ActionClusterNames(),
	})

	createCmd.AddCommand(create_clusterCmd)
}
