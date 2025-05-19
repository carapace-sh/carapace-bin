package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tsh"
	"github.com/spf13/cobra"
)

var proxy_sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "Start local TLS proxy for ssh connections when using Teleport in single-port mode.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proxy_sshCmd).Standalone()

	proxy_sshCmd.Flags().StringP("cluster", "c", "", "Specify the Teleport cluster to connect")
	proxyCmd.AddCommand(proxy_sshCmd)

	carapace.Gen(proxy_sshCmd).FlagCompletion(carapace.ActionMap{
		"cluster": tsh.ActionClusters(),
	})
}
