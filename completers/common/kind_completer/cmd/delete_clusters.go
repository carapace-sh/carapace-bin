package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var delete_clustersCmd = &cobra.Command{
	Use:   "clusters",
	Short: "Deletes one or more clusters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delete_clustersCmd).Standalone()

	delete_clustersCmd.Flags().BoolP("all", "A", false, "delete all clusters")
	delete_clustersCmd.Flags().String("kubeconfig", "", "sets kubeconfig path instead of $KUBECONFIG or $HOME/.kube/config")

	carapace.Gen(delete_clustersCmd).FlagCompletion(carapace.ActionMap{
		"kubeconfig": carapace.ActionFiles(),
	})

	deleteCmd.AddCommand(delete_clustersCmd)
}
