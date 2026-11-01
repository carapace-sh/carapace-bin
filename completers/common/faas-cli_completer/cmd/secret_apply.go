package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/faas-cli_completer/cmd/action"
	"github.com/spf13/cobra"
)

var secret_applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply secrets from .secrets folder",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_applyCmd).Standalone()

	secret_applyCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	secret_applyCmd.Flags().StringP("namespace", "n", "", "Namespace of the function")
	secret_applyCmd.Flags().Bool("tls-no-verify", false, "Disable TLS validation")
	secret_applyCmd.Flags().StringP("token", "k", "", "Pass a JWT token to use instead of basic auth")
	secret_applyCmd.Flags().Bool("trim", true, "Trim whitespace from the start and end of the secret value")
	secretCmd.AddCommand(secret_applyCmd)

	carapace.Gen(secret_applyCmd).FlagCompletion(carapace.ActionMap{
		"namespace": action.ActionNamespaces(),
	})
}
