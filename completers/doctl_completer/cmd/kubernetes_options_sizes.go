package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var kubernetes_options_sizesCmd = &cobra.Command{
	Use:   "sizes",
	Short: "List machine sizes that can be used in a DigitalOcean Kubernetes cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kubernetes_options_sizesCmd).Standalone()
	kubernetes_optionsCmd.AddCommand(kubernetes_options_sizesCmd)
}
