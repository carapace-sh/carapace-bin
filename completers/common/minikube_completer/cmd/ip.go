package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Retrieves the IP address of the specified node",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ipCmd).Standalone()

	ipCmd.Flags().StringP("node", "n", "", "The node to get IP. Defaults to the primary control plane.")
	rootCmd.AddCommand(ipCmd)

	carapace.Gen(ipCmd).FlagCompletion(carapace.ActionMap{
		"node": action.ActionNodes(),
	})
}
