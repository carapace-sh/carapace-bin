package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var create_secret_genericCmd = &cobra.Command{
	Use:   "generic NAME [--type=string] [--from-file=[key=]source] [--from-literal=key1=value1] [--dry-run=server|client|none]",
	Short: "Create a secret from a local file, directory, or literal value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_secret_genericCmd).Standalone()

	create_secret_genericCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	create_secret_genericCmd.Flags().Bool("append-hash", false, "Append a hash of the secret to its name.")
	create_secret_genericCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	create_secret_genericCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	create_secret_genericCmd.Flags().StringSlice("from-env-file", nil, "Specify the path to a file to read lines of key=val pairs to create a secret.")
	create_secret_genericCmd.Flags().StringSlice("from-file", nil, "Key files can be specified using their file path, in which case a default name will be given to them, or optionally with a name and file path, in which case the given name will be used.  Specifying a directory will iterate each named file in the directory that is a valid secret key.")
	create_secret_genericCmd.Flags().StringSlice("from-literal", nil, "Specify a key and literal value to insert in secret (i.e. mykey=somevalue)")
	create_secret_genericCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, kyaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	create_secret_genericCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	create_secret_genericCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	create_secret_genericCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	create_secret_genericCmd.Flags().String("type", "", "The type of secret to create")
	create_secret_genericCmd.Flags().String("validate", "", "Must be one of: strict (or true), warn, ignore (or false). \"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.")
	create_secret_genericCmd.Flag("dry-run").NoOptDefVal = " "
	create_secret_genericCmd.Flag("validate").NoOptDefVal = " "
	create_secretCmd.AddCommand(create_secret_genericCmd)

	carapace.Gen(create_secret_genericCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":       kubectl.ActionDryRunModes(),
		"from-env-file": carapace.ActionFiles(),
		"from-file":     carapace.ActionFiles(),
		"output":        kubectl.ActionOutputFormats(),
		"template":      carapace.ActionFiles(),
		"validate":      kubectl.ActionValidationModes(),
	})
}
