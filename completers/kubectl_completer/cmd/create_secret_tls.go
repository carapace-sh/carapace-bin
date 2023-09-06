package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var create_secret_tlsCmd = &cobra.Command{
	Use:   "tls NAME --cert=path/to/cert/file --key=path/to/key/file [--dry-run=server|client|none]",
	Short: "Create a TLS secret",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_secret_tlsCmd).Standalone()

	create_secret_tlsCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	create_secret_tlsCmd.Flags().Bool("append-hash", false, "Append a hash of the secret to its name.")
	create_secret_tlsCmd.Flags().String("cert", "", "Path to PEM encoded public key certificate.")
	create_secret_tlsCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	create_secret_tlsCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	create_secret_tlsCmd.Flags().String("key", "", "Path to private key associated with given certificate.")
	create_secret_tlsCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	create_secret_tlsCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	create_secret_tlsCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	create_secret_tlsCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	create_secret_tlsCmd.Flags().String("validate", "", "Validation mode.")
	create_secret_tlsCmd.Flag("dry-run").NoOptDefVal = " "
	create_secret_tlsCmd.Flag("validate").NoOptDefVal = " "
	create_secretCmd.AddCommand(create_secret_tlsCmd)

	carapace.Gen(create_secret_tlsCmd).FlagCompletion(carapace.ActionMap{
		"cert":     carapace.ActionFiles(),
		"dry-run":  kubectl.ActionDryRunModes(),
		"key":      carapace.ActionFiles(),
		"output":   kubectl.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
		"validate": kubectl.ActionValidationModes(),
	})
}
