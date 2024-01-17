package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var create_rolebindingCmd = &cobra.Command{
	Use:   "rolebinding NAME --clusterrole=NAME|--role=NAME [--user=username] [--group=groupname] [--serviceaccount=namespace:serviceaccountname] [--dry-run=server|client|none]",
	Short: "Create a role binding for a particular role or cluster role",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_rolebindingCmd).Standalone()

	create_rolebindingCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	create_rolebindingCmd.Flags().String("clusterrole", "", "ClusterRole this RoleBinding should reference")
	create_rolebindingCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	create_rolebindingCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	create_rolebindingCmd.Flags().StringSlice("group", []string{}, "Groups to bind to the role. The flag can be repeated to add multiple groups.")
	create_rolebindingCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	create_rolebindingCmd.Flags().String("role", "", "Role this RoleBinding should reference")
	create_rolebindingCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	create_rolebindingCmd.Flags().StringSlice("serviceaccount", []string{}, "Service accounts to bind to the role, in the format <namespace>:<name>. The flag can be repeated to add multiple service accounts.")
	create_rolebindingCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	create_rolebindingCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	create_rolebindingCmd.Flags().StringSlice("user", []string{}, "Usernames to bind to the role. The flag can be repeated to add multiple users.")
	create_rolebindingCmd.Flags().String("validate", "", "Validation mode.")
	create_rolebindingCmd.Flag("dry-run").NoOptDefVal = " "
	create_rolebindingCmd.Flag("validate").NoOptDefVal = " "
	createCmd.AddCommand(create_rolebindingCmd)

	carapace.Gen(create_rolebindingCmd).FlagCompletion(carapace.ActionMap{
		"clusterrole": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return kubectl.ActionResources(kubectl.ResourceOpts{
				Namespace: rootCmd.Flag("namespace").Value.String(),
				Types:     "clusterroles",
			})
		}),
		"dry-run": kubectl.ActionDryRunModes(),
		"output":  kubectl.ActionOutputFormats(),
		"role": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return kubectl.ActionResources(kubectl.ResourceOpts{
				Namespace: rootCmd.Flag("namespace").Value.String(),
				Types:     "roles",
			})
		}),
		"serviceaccount": kubectl.ActionNamespaceServiceAccounts(),
		"template":       carapace.ActionFiles(),
		"validate":       kubectl.ActionValidationModes(),
	})
}
