package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var secret_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all secrets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_listCmd).Standalone()
	secret_listCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	secret_listCmd.Flags().StringP("namespace", "n", "", "Namespace of the function")
	secret_listCmd.Flags().Bool("tls-no-verify", false, "Disable TLS validation")
	secret_listCmd.Flags().StringP("token", "k", "", "Pass a JWT token to use instead of basic auth")
	secretCmd.AddCommand(secret_listCmd)

	carapace.Gen(secret_listCmd).FlagCompletion(carapace.ActionMap{
		"namespace": action.ActionNamespaces(),
	})
}
