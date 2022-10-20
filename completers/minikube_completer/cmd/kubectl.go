package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var kubectlCmd = &cobra.Command{
	Use:   "kubectl",
	Short: "Run a kubectl binary matching the cluster version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	kubectlCmd.Flags().Bool("ssh", false, "Use SSH for running kubernetes client on the node")
	rootCmd.AddCommand(kubectlCmd)

	carapace.Gen(kubectlCmd).DashAnyCompletion(
		bridge.ActionCarapaceBin("kubectl"),
	)
}
