package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var create_quotaCmd = &cobra.Command{
	Use:     "quota NAME [--hard=key1=value1,key2=value2] [--scopes=Scope1,Scope2] [--dry-run=server|client|none]",
	Short:   "Create a quota with the specified name",
	Aliases: []string{"resourcequota"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_quotaCmd).Standalone()

	create_quotaCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	create_quotaCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	create_quotaCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	create_quotaCmd.Flags().String("hard", "", "A comma-delimited set of resource=quantity pairs that define a hard limit.")
	create_quotaCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	create_quotaCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	create_quotaCmd.Flags().String("scopes", "", "A comma-delimited set of quota scopes that must all match each object tracked by the quota.")
	create_quotaCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	create_quotaCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	create_quotaCmd.Flags().String("validate", "", "Must be one of: strict (or true), warn, ignore (or false). \"true\" or \"strict\" will use a schema to validate the input and fail the request if invalid. It will perform server side validation if ServerSideFieldValidation is enabled on the api-server, but will fall back to less reliable client-side validation if not. \"warn\" will warn about unknown or duplicate fields without blocking the request if server-side field validation is enabled on the API server, and behave as \"ignore\" otherwise. \"false\" or \"ignore\" will not perform any schema validation, silently dropping any unknown or duplicate fields.")
	create_quotaCmd.Flag("dry-run").NoOptDefVal = " "
	create_quotaCmd.Flag("validate").NoOptDefVal = " "
	createCmd.AddCommand(create_quotaCmd)

	carapace.Gen(create_quotaCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  kubectl.ActionDryRunModes(),
		"output":   kubectl.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
		"validate": kubectl.ActionValidationModes(),
	})
}
