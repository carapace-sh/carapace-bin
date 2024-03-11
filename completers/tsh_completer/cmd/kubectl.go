package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var kubectlCmd = &cobra.Command{
	Use:   "kubectl",
	Short: "Runs a kubectl command on a Kubernetes cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubectlCmd).Standalone()

	rootCmd.AddCommand(kubectlCmd)

	kubectlCmd.Flags().SetInterspersed(false)

	// TODO proxy the kubectl command
	carapace.Gen(kubectlCmd).PositionalAnyCompletion(
		bridge.ActionCarapaceBin("kubectl"),
	)
}
