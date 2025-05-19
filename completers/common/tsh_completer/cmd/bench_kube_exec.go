package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bench_kube_execCmd = &cobra.Command{
	Use:    "exec",
	Short:  "Run a benchmark test to exec into the specified Pod",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bench_kube_execCmd).Standalone()

	bench_kube_execCmd.Flags().String("container", "", "Selects the container to exec into.")
	bench_kube_execCmd.Flags().Bool("interactive", false, "Create interactive Kube session")
	bench_kube_execCmd.Flags().Bool("no-interactive", false, "Create interactive Kube session")
	bench_kube_execCmd.Flag("no-interactive").Hidden = true
	bench_kubeCmd.AddCommand(bench_kube_execCmd)
}
