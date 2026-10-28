package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var namespaceCmd = &cobra.Command{
	Use:     "namespace",
	Short:   "Manage OpenFaaS namespaces",
	Aliases: []string{"ns"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(namespaceCmd).Standalone()

	namespaceCmd.PersistentFlags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	namespaceCmd.PersistentFlags().Bool("tls-no-verify", false, "Disable TLS validation")
	namespaceCmd.PersistentFlags().StringP("token", "k", "", "Pass a JWT token to use instead of basic auth")
	rootCmd.AddCommand(namespaceCmd)
}
