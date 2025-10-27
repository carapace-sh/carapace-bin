package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/minikube_completer/cmd/action"
	"github.com/spf13/cobra"
)

var serviceCmd = &cobra.Command{
	Use:     "service [flags] SERVICE",
	Short:   "Returns a URL to connect to a service",
	GroupID: "networking",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceCmd).Standalone()

	serviceCmd.Flags().Bool("all", false, "Forwards all services in a namespace (defaults to \"false\")")
	serviceCmd.PersistentFlags().String("format", "", "Format to output service URL in. This format will be applied to each url individually and they will be printed one at a time.")
	serviceCmd.Flags().Bool("https", false, "Open the service URL with https instead of http (defaults to \"false\")")
	serviceCmd.Flags().String("interval", "", "The initial time interval for each check that wait performs in seconds")
	serviceCmd.Flags().StringP("namespace", "n", "", "The service namespace")
	serviceCmd.Flags().Bool("url", false, "Display the Kubernetes service URL in the CLI instead of opening it in the default browser")
	serviceCmd.Flags().String("wait", "", "Amount of time to wait for a service in seconds")
	rootCmd.AddCommand(serviceCmd)

	carapace.Gen(serviceCmd).FlagCompletion(carapace.ActionMap{
		"namespace": action.ActionNamespaces(),
	})
}
