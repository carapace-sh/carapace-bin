package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/k3d"
	"github.com/spf13/cobra"
)

var kubeconfig_getCmd = &cobra.Command{
	Use:     "get [CLUSTER [CLUSTER [...]] | --all]",
	Short:   "Print kubeconfig(s) from cluster(s).",
	Aliases: []string{"print", "show"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubeconfig_getCmd).Standalone()

	kubeconfig_getCmd.Flags().BoolP("all", "a", false, "Output kubeconfigs from all existing clusters")
	kubeconfigCmd.AddCommand(kubeconfig_getCmd)

	carapace.Gen(kubeconfig_getCmd).PositionalAnyCompletion(
		k3d.ActionClusters().FilterArgs(),
	)
}
