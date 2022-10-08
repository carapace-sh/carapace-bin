package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var create_rolebindingCmd = &cobra.Command{
	Use:   "rolebinding",
	Short: "Create a role binding for a particular role or cluster role",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_rolebindingCmd).Standalone()
	create_rolebindingCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	create_rolebindingCmd.Flags().String("clusterrole", "", "ClusterRole this RoleBinding should reference")
	create_rolebindingCmd.Flags().String("dry-run", "none", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	create_rolebindingCmd.Flags().String("field-manager", "kubectl-create", "Name of the manager used to track field ownership.")
	create_rolebindingCmd.Flags().StringArray("group", []string{}, "Groups to bind to the role. The flag can be repeated to add multiple groups.")
	create_rolebindingCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	create_rolebindingCmd.Flags().String("role", "", "Role this RoleBinding should reference")
	create_rolebindingCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	create_rolebindingCmd.Flags().StringArray("serviceaccount", []string{}, "Service accounts to bind to the role, in the format <namespace>:<name>. The flag can be repeated to add multiple service accounts.")
	create_rolebindingCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	create_rolebindingCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	create_rolebindingCmd.Flags().String("validate", "strict", "Must be one of: strict (or true), warn, ignore (or false).")
	create_rolebindingCmd.Flag("dry-run").NoOptDefVal = "unchanged"
	create_rolebindingCmd.Flag("validate").NoOptDefVal = "strict"
	createCmd.AddCommand(create_rolebindingCmd)

	carapace.Gen(create_rolebindingCmd).FlagCompletion(carapace.ActionMap{
		"clusterrole":    action.ActionResources("", "clusterroles"),
		"dry-run":        action.ActionDryRunModes(),
		"output":         action.ActionOutputFormats(),
		"role":           action.ActionResources("", "roles"),
		"serviceaccount": action.ActionNamespaceServiceAccounts(),
		"template":       carapace.ActionFiles(),
		"validate":       kubectl.ActionValidationModes(),
	})
}
