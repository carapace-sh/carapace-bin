package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var apply_editLastAppliedCmd = &cobra.Command{
	Use:   "edit-last-applied (RESOURCE/NAME | -f FILENAME)",
	Short: "Edit latest last-applied-configuration annotations of a resource/object",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apply_editLastAppliedCmd).Standalone()

	apply_editLastAppliedCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	apply_editLastAppliedCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	apply_editLastAppliedCmd.Flags().StringSliceP("filename", "f", []string{}, "Filename, directory, or URL to files to use to edit the resource")
	apply_editLastAppliedCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	apply_editLastAppliedCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	apply_editLastAppliedCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists.")
	apply_editLastAppliedCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	apply_editLastAppliedCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	apply_editLastAppliedCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	apply_editLastAppliedCmd.Flags().String("validate", "", "Validation mode.")
	apply_editLastAppliedCmd.Flags().Bool("windows-line-endings", false, "Defaults to the line ending native to your platform.")
	apply_editLastAppliedCmd.Flag("record").Hidden = true
	apply_editLastAppliedCmd.Flag("validate").NoOptDefVal = " "
	applyCmd.AddCommand(apply_editLastAppliedCmd)

	carapace.Gen(apply_editLastAppliedCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    kubectl.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
		"validate":  kubectl.ActionValidationModes(),
	})

	carapace.Gen(apply_editLastAppliedCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if apply_editLastAppliedCmd.Flag("filename").Changed {
				return carapace.ActionValues()
			} else {
				return kubectl.ActionApiResourceResources()
			}
		}),
	)
}
