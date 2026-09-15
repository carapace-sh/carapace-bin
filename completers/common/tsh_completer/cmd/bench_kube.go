package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bench_kubeCmd = &cobra.Command{
	Use:    "kube",
	Short:  "Run Kube benchmark tests.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bench_kubeCmd).Standalone()

	bench_kubeCmd.Flags().String("kube-namespace", "default", "Selects the Kubernetes namespace.")
	bench_kubeCmd.Flags().String("namespace", "default", "Selects the Kubernetes namespace.")
	bench_kubeCmd.Flag("kube-namespace").Hidden = true
	benchCmd.AddCommand(bench_kubeCmd)
}
