package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/k3d"
	"github.com/spf13/cobra"
)

var kubeconfig_mergeCmd = &cobra.Command{
	Use:     "merge [CLUSTER [CLUSTER [...]] | --all]",
	Short:   "Write/Merge kubeconfig(s) from cluster(s) into new or existing kubeconfig/file.",
	Aliases: []string{"write"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubeconfig_mergeCmd).Standalone()

	kubeconfig_mergeCmd.Flags().BoolP("all", "a", false, "Get kubeconfigs from all existing clusters")
	kubeconfig_mergeCmd.Flags().BoolP("kubeconfig-merge-default", "d", false, "Merge into the default kubeconfig ($KUBECONFIG or /root/.kube/config)")
	kubeconfig_mergeCmd.Flags().BoolP("kubeconfig-switch-context", "s", false, "Switch to new context")
	kubeconfig_mergeCmd.Flags().StringP("output", "o", "", "Define output [ - | FILE ] (default from $KUBECONFIG or /root/.kube/config")
	kubeconfig_mergeCmd.Flags().Bool("overwrite", false, "[Careful!] Overwrite existing file, ignoring its contents")
	kubeconfig_mergeCmd.Flags().BoolP("update", "u", false, "Update conflicting fields in existing kubeconfig")
	kubeconfigCmd.AddCommand(kubeconfig_mergeCmd)

	carapace.Gen(kubeconfig_mergeCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})

	carapace.Gen(kubeconfig_mergeCmd).PositionalAnyCompletion(
		k3d.ActionClusters().FilterArgs(),
	)
}
