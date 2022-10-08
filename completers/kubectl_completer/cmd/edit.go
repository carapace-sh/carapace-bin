package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a resource on the server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(editCmd).Standalone()
	editCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	editCmd.Flags().String("field-manager", "kubectl-edit", "Name of the manager used to track field ownership.")
	editCmd.Flags().StringSliceP("filename", "f", []string{}, "Filename, directory, or URL to files to use to edit the resource")
	editCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	editCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	editCmd.Flags().Bool("output-patch", false, "Output the patch if the resource is edited.")
	editCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	editCmd.Flags().Bool("save-config", false, "If true, the configuration of current object will be saved in its annotation. Otherwise, the annotation will be unchanged. This flag is useful when you want to perform kubectl apply on this object in the future.")
	editCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	editCmd.Flags().String("subresource", "", "If specified, edit will operate on the subresource of the requested object. Must be one of [status]. This flag is alpha and may change in the future.")
	editCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	editCmd.Flags().String("validate", "strict", "Must be one of: strict (or true), warn, ignore (or false).")
	editCmd.Flags().Bool("windows-line-endings", false, "Defaults to the line ending native to your platform.")
	editCmd.Flag("validate").NoOptDefVal = "strict"
	rootCmd.AddCommand(editCmd)

	carapace.Gen(editCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    action.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
		"validate":  kubectl.ActionValidationModes(),
	})

	carapace.Gen(editCmd).PositionalCompletion(
		action.ActionApiResourceResources(),
	)
}
