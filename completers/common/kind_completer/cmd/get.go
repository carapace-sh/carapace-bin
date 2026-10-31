package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Gets one of [clusters, nodes, kubeconfig]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	rootCmd.AddCommand(getCmd)
}
