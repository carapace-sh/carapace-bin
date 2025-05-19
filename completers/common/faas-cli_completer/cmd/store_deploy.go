package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/faas-cli_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/spf13/cobra"
)

var store_deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy OpenFaaS functions from a store",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_deployCmd).Standalone()
	store_deployCmd.Flags().StringArray("annotation", nil, "Set one or more annotation (ANNOTATION=VALUE)")
	store_deployCmd.Flags().StringArray("constraint", nil, "Apply a constraint to the function")
	store_deployCmd.Flags().StringArrayP("env", "e", nil, "Adds one or more environment variables to the defined ones by store (ENVVAR=VALUE)")
	store_deployCmd.Flags().StringP("gateway", "g", "http://127.0.0.1:8080", "Gateway URL starting with http(s)://")
	store_deployCmd.Flags().StringArrayP("label", "l", nil, "Set one or more label (LABEL=VALUE)")
	store_deployCmd.Flags().String("name", "", "Name of the deployed function (overriding name from the store)")
	store_deployCmd.Flags().StringP("namespace", "n", "", "Namespace of the function")
	store_deployCmd.Flags().String("network", "", "Name of the network")
	store_deployCmd.Flags().Bool("replace", false, "Replace any existing function")
	store_deployCmd.Flags().StringArray("secret", nil, "Give the function access to a secure secret")
	store_deployCmd.Flags().Bool("tls-no-verify", false, "Disable TLS validation")
	store_deployCmd.Flags().StringP("token", "k", "", "Pass a JWT token to use instead of basic auth")
	store_deployCmd.Flags().Bool("update", true, "Update existing functions")
	storeCmd.AddCommand(store_deployCmd)

	carapace.Gen(store_deployCmd).FlagCompletion(carapace.ActionMap{
		"env":       env.ActionNameValues(false),
		"namespace": action.ActionNamespaces(),
	})

	carapace.Gen(store_deployCmd).PositionalCompletion(
		action.ActionFunctions(),
	)
}
