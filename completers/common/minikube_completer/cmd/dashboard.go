package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var dashboardCmd = &cobra.Command{
	Use:     "dashboard",
	Short:   "Access the Kubernetes dashboard running within the minikube cluster",
	GroupID: "basic",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dashboardCmd).Standalone()

	dashboardCmd.Flags().String("port", "", "Exposed port of the proxyfied dashboard. Set to 0 to pick a random port.")
	dashboardCmd.Flags().Bool("url", false, "Display dashboard URL instead of opening a browser")
	rootCmd.AddCommand(dashboardCmd)

	carapace.Gen(dashboardCmd).FlagCompletion(carapace.ActionMap{
		"port": net.ActionPorts(),
	})
}
