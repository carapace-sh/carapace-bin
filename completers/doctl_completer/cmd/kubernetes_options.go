package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_optionsCmd = &cobra.Command{
	Use:   "options",
	Short: "List possible option values for use inside Kubernetes commands",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_optionsCmd).Standalone()
	kubernetesCmd.AddCommand(kubernetes_optionsCmd)
}
