package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_options_regionsCmd = &cobra.Command{
	Use:   "regions",
	Short: "List regions that support DigitalOcean Kubernetes clusters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_options_regionsCmd).Standalone()
	kubernetes_optionsCmd.AddCommand(kubernetes_options_regionsCmd)
}
