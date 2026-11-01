package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/faas-cli_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Kubernetes CRD YAML file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateCmd).Standalone()

	generateCmd.Flags().StringArray("annotation", nil, "Any annotations you want to add (to store functions only)")
	generateCmd.Flags().String("api", "openfaas.com/v1", "CRD API version e.g openfaas.com/v1")
	generateCmd.Flags().String("arch", "x86_64", "Desired image arch. (Default x86_64)")
	generateCmd.Flags().StringArrayP("env", "e", nil, "Set one or more environment variables (ENVVAR=VALUE) (to store functions only)")
	generateCmd.Flags().Bool("envsubst", true, "Substitute environment variables in stack.yaml file")
	generateCmd.Flags().String("from-store", "", "generate using a store image")
	generateCmd.Flags().StringArray("label", nil, "Any labels you want to add (to store functions only)")
	generateCmd.Flags().String("name", "", "for use with --from-store, override the name for the generated function")
	generateCmd.Flags().StringP("namespace", "n", "openfaas-fn", "Kubernetes namespace for functions")
	generateCmd.Flags().String("output", "", "output format e.g stack.yaml, for use with --from-store to generate an OpenFaaS stack.yaml")
	generateCmd.Flags().String("tag", "latest", "Override latest tag on function Docker image, accepts 'digest', 'latest', 'sha', 'branch', 'describe'")
	rootCmd.AddCommand(generateCmd)

	carapace.Gen(generateCmd).FlagCompletion(carapace.ActionMap{
		"api":        carapace.ActionValues("openfaas.com/v1"),
		"arch":       action.ActionArchitectures(),
		"env":        env.ActionNameValues(false),
		"from-store": action.ActionFunctionsStore(),
		"namespace":  action.ActionNamespaces(),
		"tag":        carapace.ActionValues("branch", "describe", "digest", "latest", "sha"),
	})
}
