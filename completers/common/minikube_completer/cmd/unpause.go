package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var unpauseCmd = &cobra.Command{
	Use:     "unpause",
	Short:   "unpause Kubernetes",
	GroupID: "basic",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unpauseCmd).Standalone()
	unpauseCmd.Flags().BoolP("all-namespaces", "A", false, "If set, unpause all namespaces")
	unpauseCmd.Flags().StringSliceP("namespaces", "n", []string{"kube-system", "kubernetes-dashboard", "storage-gluster", "istio-operator"}, "namespaces to unpause")
	unpauseCmd.Flags().StringP("output", "o", "text", "Format to print stdout in. Options include: [text,json]")
	rootCmd.AddCommand(unpauseCmd)

	carapace.Gen(unpauseCmd).FlagCompletion(carapace.ActionMap{
		"namespaces": action.ActionNamespaces().UniqueList(","),
		"output":     carapace.ActionValues("text", "json"),
	})
}
