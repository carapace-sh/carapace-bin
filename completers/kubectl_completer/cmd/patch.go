package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var patchCmd = &cobra.Command{
	Use:     "patch (-f FILENAME | TYPE NAME) [-p PATCH|--patch-file FILE]",
	Short:   "Update fields of a resource",
	GroupID: "advanced",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(patchCmd).Standalone()

	patchCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	patchCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	patchCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	patchCmd.Flags().StringSliceP("filename", "f", nil, "Filename, directory, or URL to files identifying the resource to update")
	patchCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	patchCmd.Flags().Bool("local", false, "If true, patch will operate on the content of the file, not the server-side resource.")
	patchCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	patchCmd.Flags().StringP("patch", "p", "", "The patch to be applied to the resource JSON file.")
	patchCmd.Flags().String("patch-file", "", "A file containing a patch to be applied to the resource.")
	patchCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists.")
	patchCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	patchCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	patchCmd.Flags().String("subresource", "", "If specified, patch will operate on the subresource of the requested object.")
	patchCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	patchCmd.Flags().String("type", "", "The type of patch being provided; one of [json merge strategic]")
	patchCmd.Flag("dry-run").NoOptDefVal = " "
	patchCmd.Flag("record").Hidden = true
	rootCmd.AddCommand(patchCmd)

	carapace.Gen(patchCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":    kubectl.ActionDryRunModes(),
		"filename":   carapace.ActionFiles(),
		"kustomize":  carapace.ActionDirectories(),
		"output":     kubectl.ActionOutputFormats(),
		"patch-file": carapace.ActionFiles(),
		"template":   carapace.ActionFiles(),
		"type":       carapace.ActionValues("json", "merge", "strategic"),
	})

	carapace.Gen(patchCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if patchCmd.Flag("filename").Changed {
				return carapace.ActionValues()
			} else {
				return kubectl.ActionApiResources()
			}
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if patchCmd.Flag("filename").Changed {
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
