package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var set_subjectCmd = &cobra.Command{
	Use:   "subject (-f FILENAME | TYPE NAME) [--user=username] [--group=groupname] [--serviceaccount=namespace:serviceaccountname] [--dry-run=server|client|none]",
	Short: "Update the user, group, or service account in a role binding or cluster role binding",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_subjectCmd).Standalone()

	set_subjectCmd.Flags().Bool("all", false, "Select all resources, in the namespace of the specified resource types")
	set_subjectCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	set_subjectCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	set_subjectCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	set_subjectCmd.Flags().StringSliceP("filename", "f", nil, "Filename, directory, or URL to files the resource to update the subjects")
	set_subjectCmd.Flags().StringSlice("group", nil, "Groups to bind to the role")
	set_subjectCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	set_subjectCmd.Flags().Bool("local", false, "If true, set subject will NOT contact api-server but run locally.")
	set_subjectCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	set_subjectCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	set_subjectCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.")
	set_subjectCmd.Flags().StringSlice("serviceaccount", nil, "Service accounts to bind to the role")
	set_subjectCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	set_subjectCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	set_subjectCmd.Flags().StringSlice("user", nil, "Usernames to bind to the role")
	set_subjectCmd.Flag("dry-run").NoOptDefVal = " "
	setCmd.AddCommand(set_subjectCmd)

	carapace.Gen(set_subjectCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":        kubectl.ActionDryRunModes(),
		"filename":       carapace.ActionFiles(),
		"kustomize":      carapace.ActionDirectories(),
		"output":         kubectl.ActionOutputFormats(),
		"serviceaccount": kubectl.ActionNamespaceServiceAccounts(),
		"template":       carapace.ActionFiles(),
	})

	carapace.Gen(set_subjectCmd).PositionalCompletion(
		carapace.ActionValues("rolebinding", "clusterrolebinding"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return kubectl.ActionResources(kubectl.ResourceOpts{
				Context:   rootCmd.Flag("context").Value.String(),
				Namespace: rootCmd.Flag("namespace").Value.String(),
				Types:     c.Args[0],
			})
		}),
	)
}
