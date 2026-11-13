package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/kind_completer/cmd/action"
	"github.com/spf13/cobra"
)

var get_kubeconfigCmd = &cobra.Command{
	Use:   "kubeconfig",
	Short: "Prints cluster kubeconfig",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_kubeconfigCmd).Standalone()

	get_kubeconfigCmd.Flags().Bool("internal", false, "use internal address instead of external")
	get_kubeconfigCmd.Flags().StringP("name", "n", "kind", "the cluster context name")

	carapace.Gen(get_kubeconfigCmd).FlagCompletion(carapace.ActionMap{
		"name": action.ActionClusterNames(),
	})

	getCmd.AddCommand(get_kubeconfigCmd)
}
