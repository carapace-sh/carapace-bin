package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetesCmd = &cobra.Command{
	Use:   "kubernetes",
	Short: "Displays commands to manage Kubernetes clusters and configurations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetesCmd).Standalone()
	rootCmd.AddCommand(kubernetesCmd)
}
