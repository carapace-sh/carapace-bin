package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kube_downCmd = &cobra.Command{
	Use:   "down [options] KUBEFILE|-",
	Short: "Remove pods based on Kubernetes YAML",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kube_downCmd).Standalone()

	kube_downCmd.Flags().Bool("force", false, "remove volumes")
	kubeCmd.AddCommand(kube_downCmd)
}
