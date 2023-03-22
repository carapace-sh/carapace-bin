package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pauseCmd = &cobra.Command{
	Use:     "pause",
	Short:   "pause Kubernetes",
	GroupID: "basic",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	pauseCmd.Flags().BoolP("all-namespaces", "A", false, "If set, pause all namespaces")
	pauseCmd.Flags().StringSliceP("namespaces", "n", []string{"kube-system", "kubernetes-dashboard", "storage-gluster", "istio-operator"}, "namespaces to pause")
	pauseCmd.Flags().StringP("output", "o", "text", "Format to print stdout in. Options include: [text,json]")
	rootCmd.AddCommand(pauseCmd)

	carapace.Gen(pauseCmd).FlagCompletion(carapace.ActionMap{
		"namespaces": action.ActionNamespaces().UniqueList(","),
		"output":     carapace.ActionValues("text", "json"),
	})
}
