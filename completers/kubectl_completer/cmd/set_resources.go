package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var set_resourcesCmd = &cobra.Command{
	Use:   "resources (-f FILENAME | TYPE NAME)  ([--limits=LIMITS & --requests=REQUESTS]",
	Short: "Update resource requests/limits on objects with pod templates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_resourcesCmd).Standalone()

	set_resourcesCmd.Flags().Bool("all", false, "Select all resources, in the namespace of the specified resource types")
	set_resourcesCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	set_resourcesCmd.Flags().StringP("containers", "c", "", "The names of containers in the selected pod templates to change, all containers are selected by default - may use wildcards")
	set_resourcesCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	set_resourcesCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	set_resourcesCmd.Flags().StringSliceP("filename", "f", nil, "Filename, directory, or URL to files identifying the resource to get from a server.")
	set_resourcesCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	set_resourcesCmd.Flags().String("limits", "", "The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'.  Note that server side components may assign requests depending on the server configuration, such as limit ranges.")
	set_resourcesCmd.Flags().Bool("local", false, "If true, set resources will NOT contact api-server but run locally.")
	set_resourcesCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	set_resourcesCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists.")
	set_resourcesCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	set_resourcesCmd.Flags().String("requests", "", "The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'.  Note that server side components may assign requests depending on the server configuration, such as limit ranges.")
	set_resourcesCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.")
	set_resourcesCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	set_resourcesCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	set_resourcesCmd.Flag("dry-run").NoOptDefVal = " "
	set_resourcesCmd.Flag("record").Hidden = true
	setCmd.AddCommand(set_resourcesCmd)

	carapace.Gen(set_resourcesCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   kubectl.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    kubectl.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(set_resourcesCmd).PositionalCompletion(
		carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if set_resourcesCmd.Flag("filename").Changed {
				return carapace.ActionValues()
			} else {
				return kubectl.ActionApiResources().NoSpace()
			}
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if set_resourcesCmd.Flag("filename").Changed {
				return carapace.ActionValues()
			} else {
				return kubectl.ActionResources(kubectl.ResourceOpts{
					Context:   rootCmd.Flag("context").Value.String(),
					Namespace: rootCmd.Flag("namespace").Value.String(),
					Types:     c.Args[0],
				})
			}
		}),
	)
}
