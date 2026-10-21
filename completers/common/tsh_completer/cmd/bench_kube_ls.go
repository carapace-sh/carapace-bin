package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var bench_kube_lsCmd = &cobra.Command{
	Use:    "ls",
	Short:  "Run a benchmark test to list Pods.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bench_kube_lsCmd).Standalone()

	bench_kubeCmd.AddCommand(bench_kube_lsCmd)

	carapace.Gen(bench_kube_lsCmd).PositionalCompletion(
		tsh.ActionKubeClusters(),
	)
}
