package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_1ClickCmd = &cobra.Command{
	Use:   "1-click",
	Short: "Display commands that pertain to kubernetes 1-click applications",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_1ClickCmd).Standalone()
	kubernetesCmd.AddCommand(kubernetes_1ClickCmd)
}
