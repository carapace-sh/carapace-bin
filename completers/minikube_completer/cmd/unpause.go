package cmd

import (
	"github.com/spf13/cobra"
)

var unpauseCmd = &cobra.Command{
	Use:   "unpause",
	Short: "unpause Kubernetes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	unpauseCmd.Flags().StringSliceP("--namespaces", "n", []string{"kube-system", "kubernetes-dashboard", "storage-gluster", "istio-operator"}, "namespaces to unpause")
	unpauseCmd.Flags().BoolP("all-namespaces", "A", false, "If set, unpause all namespaces")
	unpauseCmd.Flags().StringP("output", "o", "text", "Format to print stdout in. Options include: [text,json]")
	rootCmd.AddCommand(unpauseCmd)
}
