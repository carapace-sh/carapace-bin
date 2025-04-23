package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var annotateCmd = &cobra.Command{
	Use:     "annotate [--overwrite] (-f FILENAME | TYPE NAME) KEY_1=VAL_1 ... KEY_N=VAL_N [--resource-version=version]",
	Short:   "Update the annotations on a resource",
	GroupID: "settings",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(annotateCmd).Standalone()

	annotateCmd.Flags().Bool("all", false, "Select all resources, in the namespace of the specified resource types.")
	annotateCmd.Flags().BoolP("all-namespaces", "A", false, "If true, check the specified action in all namespaces.")
	annotateCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	annotateCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	annotateCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	annotateCmd.Flags().String("field-selector", "", "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type.")
	annotateCmd.Flags().StringSliceP("filename", "f", []string{}, "Filename, directory, or URL to files identifying the resource to update the annotation")
	annotateCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	annotateCmd.Flags().Bool("list", false, "If true, display the annotations for a given resource.")
	annotateCmd.Flags().Bool("local", false, "If true, annotation will NOT contact api-server but run locally.")
	annotateCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	annotateCmd.Flags().Bool("overwrite", false, "If true, allow annotations to be overwritten, otherwise reject annotation updates that overwrite existing annotations.")
	annotateCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists.")
	annotateCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	annotateCmd.Flags().String("resource-version", "", "If non-empty, the annotation update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource.")
	annotateCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.")
	annotateCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	annotateCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	annotateCmd.Flag("dry-run").NoOptDefVal = " "
	annotateCmd.Flag("record").Hidden = true
	rootCmd.AddCommand(annotateCmd)

	carapace.Gen(annotateCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   kubectl.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    kubectl.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(annotateCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if labelCmd.Flag("filename").Changed {
				return carapace.ActionValues() // TODO get labels for file
			} else {
				return kubectl.ActionApiResources()
			}
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if labelCmd.Flag("all").Changed {
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

	carapace.Gen(annotateCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if labelCmd.Flag("filename").Changed || labelCmd.Flag("all").Changed {
				return carapace.ActionValues() // TODO support labels for file
			} else {
				return kubectl.ActionAnnotations(kubectl.AnnotationOpts{Namespace: "", Resource: c.Args[0] + "/" + c.Args[1]})
			}
		}),
	)
}
