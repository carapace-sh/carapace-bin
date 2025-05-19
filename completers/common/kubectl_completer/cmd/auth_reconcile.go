package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var auth_reconcileCmd = &cobra.Command{
	Use:   "reconcile -f FILENAME",
	Short: "Reconciles rules for RBAC role, role binding, cluster role, and cluster role binding objects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(auth_reconcileCmd).Standalone()

	auth_reconcileCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	auth_reconcileCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	auth_reconcileCmd.Flags().StringSliceP("filename", "f", nil, "Filename, directory, or URL to files identifying the resource to reconcile.")
	auth_reconcileCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	auth_reconcileCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	auth_reconcileCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	auth_reconcileCmd.Flags().Bool("remove-extra-permissions", false, "If true, removes extra permissions added to roles")
	auth_reconcileCmd.Flags().Bool("remove-extra-subjects", false, "If true, removes extra subjects added to rolebindings")
	auth_reconcileCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	auth_reconcileCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	auth_reconcileCmd.Flag("dry-run").NoOptDefVal = " "
	authCmd.AddCommand(auth_reconcileCmd)

	carapace.Gen(auth_reconcileCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   kubectl.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    kubectl.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})
}
