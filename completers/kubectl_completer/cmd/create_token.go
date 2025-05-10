package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var create_tokenCmd = &cobra.Command{
	Use:   "token SERVICE_ACCOUNT_NAME",
	Short: "Request a service account token",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_tokenCmd).Standalone()

	create_tokenCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	create_tokenCmd.Flags().StringSlice("audience", nil, "Audience of the requested token. If unset, defaults to requesting a token for use with the Kubernetes API server. May be repeated to request a token valid for multiple audiences.")
	create_tokenCmd.Flags().String("bound-object-kind", "", "Kind of an object to bind the token to. Supported kinds are Node, Pod, Secret. If set, --bound-object-name must be provided.")
	create_tokenCmd.Flags().String("bound-object-name", "", "Name of an object to bind the token to. The token will expire when the object is deleted. Requires --bound-object-kind.")
	create_tokenCmd.Flags().String("bound-object-uid", "", "UID of an object to bind the token to. Requires --bound-object-kind and --bound-object-name. If unset, the UID of the existing object is used.")
	create_tokenCmd.Flags().String("duration", "", "Requested lifetime of the issued token. If not set or if set to 0, the lifetime will be determined by the server automatically. The server may return a token with a longer or shorter lifetime.")
	create_tokenCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	create_tokenCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	create_tokenCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	createCmd.AddCommand(create_tokenCmd)

	carapace.Gen(create_tokenCmd).FlagCompletion(carapace.ActionMap{
		"output":   kubectl.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
	})
}
