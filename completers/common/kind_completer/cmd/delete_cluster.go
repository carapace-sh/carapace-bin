package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/kind_completer/cmd/action"
	"github.com/spf13/cobra"
)

var delete_clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Deletes a cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delete_clusterCmd).Standalone()

	delete_clusterCmd.Flags().String("kubeconfig", "", "sets kubeconfig path instead of $KUBECONFIG or $HOME/.kube/config")
	delete_clusterCmd.Flags().StringP("name", "n", "kind", "the cluster name")

	carapace.Gen(delete_clusterCmd).FlagCompletion(carapace.ActionMap{
		"kubeconfig": carapace.ActionFiles(),
		"name":       action.ActionClusterNames(),
	})

	deleteCmd.AddCommand(delete_clusterCmd)
}
