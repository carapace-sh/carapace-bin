package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var proxy_gcloudCmd = &cobra.Command{
	Use:     "gcloud",
	Short:   "Start local proxy for GCP access.",
	Aliases: []string{"gcp"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proxy_gcloudCmd).Standalone()

	proxy_gcloudCmd.Flags().String("app", "", "Optional Name of the GCP application to use if logged into multiple.")
	proxy_gcloudCmd.Flags().StringP("format", "f", "", "Optional format to print the commands for setting environment variables, one of: unix, command-prompt, powershell, text. Default is unix.")
	proxy_gcloudCmd.Flags().StringP("port", "p", "", "Specifies the source port used by the proxy listener.")
	proxyCmd.AddCommand(proxy_gcloudCmd)

	carapace.Gen(proxy_gcloudCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("unix", "command-prompt", "powershell", "text"),
		"port":   net.ActionPorts(),
	})
}
