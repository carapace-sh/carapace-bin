package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var create_quotaCmd = &cobra.Command{
	Use:     "quota",
	Short:   "Create a quota with the specified name",
	Aliases: []string{"resourcequota"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_quotaCmd).Standalone()
	create_quotaCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	create_quotaCmd.Flags().String("dry-run", "none", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	create_quotaCmd.Flags().String("field-manager", "kubectl-create", "Name of the manager used to track field ownership.")
	create_quotaCmd.Flags().String("hard", "", "A comma-delimited set of resource=quantity pairs that define a hard limit.")
	create_quotaCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	create_quotaCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	create_quotaCmd.Flags().String("scopes", "", "A comma-delimited set of quota scopes that must all match each object tracked by the quota.")
	create_quotaCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	create_quotaCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	create_quotaCmd.Flags().String("validate", "strict", "Must be one of: strict (or true), warn, ignore (or false).")
	create_quotaCmd.Flag("dry-run").NoOptDefVal = "unchanged"
	create_quotaCmd.Flag("validate").NoOptDefVal = "strict"
	createCmd.AddCommand(create_quotaCmd)

	carapace.Gen(create_quotaCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  action.ActionDryRunModes(),
		"output":   action.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
		"validate": kubectl.ActionValidationModes(),
	})
}