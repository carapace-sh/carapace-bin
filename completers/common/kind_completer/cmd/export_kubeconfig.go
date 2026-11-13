package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/kind_completer/cmd/action"
	"github.com/spf13/cobra"
)

var export_kubeconfigCmd = &cobra.Command{
	Use:   "kubeconfig",
	Short: "Exports cluster kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(export_kubeconfigCmd).Standalone()

	export_kubeconfigCmd.Flags().Bool("internal", false, "use internal address instead of external")
	export_kubeconfigCmd.Flags().String("kubeconfig", "", "sets kubeconfig path instead of $KUBECONFIG or $HOME/.kube/config")
	export_kubeconfigCmd.Flags().StringP("name", "n", "kind", "the cluster context name")

	carapace.Gen(export_kubeconfigCmd).FlagCompletion(carapace.ActionMap{
		"kubeconfig": carapace.ActionFiles(),
		"name":       action.ActionClusterNames(),
	})

	exportCmd.AddCommand(export_kubeconfigCmd)
}
