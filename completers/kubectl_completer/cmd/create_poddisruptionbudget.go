package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var create_poddisruptionbudgetCmd = &cobra.Command{
	Use:     "poddisruptionbudget",
	Short:   "Create a pod disruption budget with the specified name",
	Aliases: []string{"pdb"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_poddisruptionbudgetCmd).Standalone()
	create_poddisruptionbudgetCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	create_poddisruptionbudgetCmd.Flags().String("dry-run", "none", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	create_poddisruptionbudgetCmd.Flags().String("field-manager", "kubectl-create", "Name of the manager used to track field ownership.")
	create_poddisruptionbudgetCmd.Flags().String("max-unavailable", "", "The maximum number or percentage of unavailable pods this budget requires.")
	create_poddisruptionbudgetCmd.Flags().String("min-available", "", "The minimum number or percentage of available pods this budget requires.")
	create_poddisruptionbudgetCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	create_poddisruptionbudgetCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	create_poddisruptionbudgetCmd.Flags().String("selector", "", "A label selector to use for this budget. Only equality-based selector requirements are supported.")
	create_poddisruptionbudgetCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	create_poddisruptionbudgetCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	create_poddisruptionbudgetCmd.Flags().String("validate", "strict", "Must be one of: strict (or true), warn, ignore (or false).")
	create_poddisruptionbudgetCmd.Flag("dry-run").NoOptDefVal = "unchanged"
	create_poddisruptionbudgetCmd.Flag("validate").NoOptDefVal = "strict"
	createCmd.AddCommand(create_poddisruptionbudgetCmd)

	carapace.Gen(create_poddisruptionbudgetCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  action.ActionDryRunModes(),
		"output":   action.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
		"validate": kubectl.ActionValidationModes(),
	})
}
