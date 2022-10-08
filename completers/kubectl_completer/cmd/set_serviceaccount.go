package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var set_serviceaccountCmd = &cobra.Command{
	Use:     "serviceaccount",
	Short:   "Update the service account of a resource",
	Aliases: []string{"sa"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_serviceaccountCmd).Standalone()
	set_serviceaccountCmd.Flags().Bool("all", false, "Select all resources, in the namespace of the specified resource types")
	set_serviceaccountCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	set_serviceaccountCmd.Flags().String("dry-run", "none", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	set_serviceaccountCmd.Flags().String("field-manager", "kubectl-set", "Name of the manager used to track field ownership.")
	set_serviceaccountCmd.Flags().StringSliceP("filename", "f", []string{}, "Filename, directory, or URL to files identifying the resource to get from a server.")
	set_serviceaccountCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	set_serviceaccountCmd.Flags().Bool("local", false, "If true, set serviceaccount will NOT contact api-server but run locally.")
	set_serviceaccountCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	set_serviceaccountCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	set_serviceaccountCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	set_serviceaccountCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	set_serviceaccountCmd.Flag("dry-run").NoOptDefVal = "unchanged"
	setCmd.AddCommand(set_serviceaccountCmd)

	carapace.Gen(set_serviceaccountCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   action.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(set_serviceaccountCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if set_serviceaccountCmd.Flag("filename").Changed {
				return action.ActionResources("", "serviceaccounts")
			} else {
				return action.ActionApiResources()
			}
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if set_serviceaccountCmd.Flag("filename").Changed {
				return carapace.ActionValues()
			} else {
				return action.ActionResources("", c.Args[0])
			}
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if set_serviceaccountCmd.Flag("filename").Changed {
				return carapace.ActionValues()
			} else {
				return action.ActionResources("", "serviceaccounts")
			}
		}),
	)
}
