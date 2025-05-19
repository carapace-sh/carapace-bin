package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var certificate_denyCmd = &cobra.Command{
	Use:   "deny (-f FILENAME | NAME)",
	Short: "Deny a certificate signing request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(certificate_denyCmd).Standalone()

	certificate_denyCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	certificate_denyCmd.Flags().StringSliceP("filename", "f", nil, "Filename, directory, or URL to files identifying the resource to update")
	certificate_denyCmd.Flags().Bool("force", false, "Update the CSR even if it is already denied.")
	certificate_denyCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	certificate_denyCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	certificate_denyCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	certificate_denyCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	certificate_denyCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	certificateCmd.AddCommand(certificate_denyCmd)

	carapace.Gen(certificate_denyCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    kubectl.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})
}
