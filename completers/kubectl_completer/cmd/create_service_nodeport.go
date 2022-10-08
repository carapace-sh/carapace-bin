package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var create_service_nodeportCmd = &cobra.Command{
	Use:   "nodeport",
	Short: "Create a NodePort service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_service_nodeportCmd).Standalone()
	create_service_nodeportCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	create_service_nodeportCmd.Flags().String("dry-run", "none", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	create_service_nodeportCmd.Flags().String("field-manager", "kubectl-create", "Name of the manager used to track field ownership.")
	create_service_nodeportCmd.Flags().Int("node-port", 0, "Port used to expose the service on each node in a cluster.")
	create_service_nodeportCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	create_service_nodeportCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	create_service_nodeportCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	create_service_nodeportCmd.Flags().StringSlice("tcp", []string{}, "Port pairs can be specified as '<port>:<targetPort>'.")
	create_service_nodeportCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	create_service_nodeportCmd.Flags().String("validate", "strict", "Must be one of: strict (or true), warn, ignore (or false).")
	create_service_nodeportCmd.Flag("dry-run").NoOptDefVal = "unchanged"
	create_service_nodeportCmd.Flag("validate").NoOptDefVal = "strict"
	create_serviceCmd.AddCommand(create_service_nodeportCmd)

	carapace.Gen(create_service_nodeportCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  action.ActionDryRunModes(),
		"output":   action.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
		"validate": kubectl.ActionValidationModes(),
	})
}
