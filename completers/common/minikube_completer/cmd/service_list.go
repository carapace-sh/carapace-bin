package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var service_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists the URLs for the services in your local cluster",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_listCmd).Standalone()
	service_listCmd.Flags().StringP("namespace", "n", "", "The services namespace")
	serviceCmd.AddCommand(service_listCmd)

	carapace.Gen(service_listCmd).FlagCompletion(carapace.ActionMap{
		"namespace": action.ActionNamespaces(),
	})
}
