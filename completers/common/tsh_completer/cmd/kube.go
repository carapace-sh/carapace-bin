package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kubeCmd = &cobra.Command{
	Use:   "kube",
	Short: "Manage available Kubernetes clusters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubeCmd).Standalone()

	rootCmd.AddCommand(kubeCmd)
}
