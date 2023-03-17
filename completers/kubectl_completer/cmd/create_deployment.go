package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var create_deploymentCmd = &cobra.Command{
	Use:     "deployment NAME --image=image -- [COMMAND] [args...]",
	Short:   "Create a deployment with the specified name",
	Aliases: []string{"deploy"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_deploymentCmd).Standalone()

	create_deploymentCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	create_deploymentCmd.Flags().String("dry-run", "none", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	create_deploymentCmd.Flags().String("field-manager", "kubectl-create", "Name of the manager used to track field ownership.")
	create_deploymentCmd.Flags().StringSlice("image", []string{}, "Image names to run.")
	create_deploymentCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	create_deploymentCmd.Flags().Int32("port", -1, "The port that this container exposes.")
	create_deploymentCmd.Flags().Int32P("replicas", "r", 1, "Number of replicas to create. Default is 1.")
	create_deploymentCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	create_deploymentCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	create_deploymentCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	create_deploymentCmd.Flags().String("validate", "strict", "Must be one of: strict (or true), warn, ignore (or false).")
	create_deploymentCmd.Flag("dry-run").NoOptDefVal = "unchanged"
	create_deploymentCmd.Flag("validate").NoOptDefVal = "strict"
	createCmd.AddCommand(create_deploymentCmd)

	carapace.Gen(create_deploymentCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  kubectl.ActionDryRunModes(),
		"output":   kubectl.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
		"validate": kubectl.ActionValidationModes(),
	})
}
