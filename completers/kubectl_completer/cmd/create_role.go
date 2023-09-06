package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var create_roleCmd = &cobra.Command{
	Use:   "role NAME --verb=verb --resource=resource.group/subresource [--resource-name=resourcename] [--dry-run=server|client|none]",
	Short: "Create a role with single rule",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_roleCmd).Standalone()

	create_roleCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	create_roleCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	create_roleCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	create_roleCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	create_roleCmd.Flags().StringSlice("resource", []string{}, "Resource that the rule applies to")
	create_roleCmd.Flags().StringSlice("resource-name", []string{}, "Resource in the white list that the rule applies to, repeat this flag for multiple items")
	create_roleCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	create_roleCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	create_roleCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	create_roleCmd.Flags().String("validate", "", "Validation mode.")
	create_roleCmd.Flags().StringSlice("verb", []string{}, "Verb that applies to the resources contained in the rule")
	create_roleCmd.Flag("dry-run").NoOptDefVal = " "
	create_roleCmd.Flag("validate").NoOptDefVal = " "
	createCmd.AddCommand(create_roleCmd)

	carapace.Gen(create_roleCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  kubectl.ActionDryRunModes(),
		"output":   kubectl.ActionOutputFormats(),
		"resource": kubectl.ActionApiResources(),
		"template": carapace.ActionFiles(),
		"validate": kubectl.ActionValidationModes(),
		"verb":     kubectl.ActionResourceVerbs(),
	})
}
