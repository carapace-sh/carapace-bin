package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var proxy_azureCmd = &cobra.Command{
	Use:     "azure",
	Short:   "Start local proxy for Azure access.",
	Aliases: []string{"az"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proxy_azureCmd).Standalone()

	proxy_azureCmd.Flags().String("app", "", "Optional Name of the Azure application to use if logged into multiple.")
	proxy_azureCmd.Flags().StringP("format", "f", "", "Optional format to print the commands for setting environment variables, one of: unix, command-prompt, powershell, text. Default is unix.")
	proxy_azureCmd.Flags().StringP("port", "p", "", "Specifies the source port used by the proxy listener.")
	proxyCmd.AddCommand(proxy_azureCmd)

	carapace.Gen(proxy_azureCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("unix", "command-prompt", "powershell", "text"),
		"port":   net.ActionPorts(),
	})
}
