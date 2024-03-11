package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var proxy_awsCmd = &cobra.Command{
	Use:   "aws",
	Short: "Start local proxy for AWS access.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proxy_awsCmd).Standalone()

	proxy_awsCmd.Flags().String("app", "", "Optional Name of the AWS application to use if logged into multiple.")
	proxy_awsCmd.Flags().BoolP("endpoint-url", "e", false, "Run local proxy to serve as an AWS endpoint URL. If not specified, local proxy serves as an HTTPS proxy.")
	proxy_awsCmd.Flags().StringP("format", "f", "", "Optional format to print the commands for setting environment variables, one of: unix, command-prompt, powershell, text. Default is unix. Or specify a service format, one of: athena-odbc, athena-jdbc")
	proxy_awsCmd.Flags().Bool("no-endpoint-url", false, "Run local proxy to serve as an AWS endpoint URL. If not specified, local proxy serves as an HTTPS proxy.")
	proxy_awsCmd.Flags().StringP("port", "p", "", "Specifies the source port used by the proxy listener.")
	proxy_awsCmd.Flag("no-endpoint-url").Hidden = true
	proxyCmd.AddCommand(proxy_awsCmd)

	carapace.Gen(proxy_awsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("unix", "command-prompt", "powershell", "text"),
		"port":   net.ActionPorts(),
	})
}
