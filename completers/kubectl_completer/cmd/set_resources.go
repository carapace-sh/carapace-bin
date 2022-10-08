package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var set_resourcesCmd = &cobra.Command{
	Use:   "resources",
	Short: "Update resource requests/limits on objects with pod templates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_resourcesCmd).Standalone()
	set_resourcesCmd.Flags().Bool("all", false, "Select all resources, in the namespace of the specified resource types")
	set_resourcesCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	set_resourcesCmd.Flags().StringP("containers", "c", "*", "The names of containers in the selected pod templates to change, all containers are selected by default - may use wildcards")
	set_resourcesCmd.Flags().String("dry-run", "none", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	set_resourcesCmd.Flags().String("field-manager", "kubectl-set", "Name of the manager used to track field ownership.")
	set_resourcesCmd.Flags().StringSliceP("filename", "f", []string{}, "Filename, directory, or URL to files identifying the resource to get from a server.")
	set_resourcesCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	set_resourcesCmd.Flags().String("limits", "", "The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'.  Note that server side components may assign requests depending on the server configuration, such as limit ranges.")
	set_resourcesCmd.Flags().Bool("local", false, "If true, set resources will NOT contact api-server but run locally.")
	set_resourcesCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	set_resourcesCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	set_resourcesCmd.Flags().String("requests", "", "The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'.  Note that server side components may assign requests depending on the server configuration, such as limit ranges.")
	set_resourcesCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2). Matching objects must satisfy all of the specified label constraints.")
	set_resourcesCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	set_resourcesCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	set_resourcesCmd.Flag("dry-run").NoOptDefVal = "unchanged"
	setCmd.AddCommand(set_resourcesCmd)

	carapace.Gen(set_resourcesCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   action.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(set_resourcesCmd).PositionalCompletion(
		carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if set_resourcesCmd.Flag("filename").Changed {
				return carapace.ActionValues()
			} else {
				return action.ActionApiResources()
			}
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if set_resourcesCmd.Flag("filename").Changed {
				return carapace.ActionValues()
			} else {
				return action.ActionResources("", c.Args[0])
			}
		}),
	)
}
