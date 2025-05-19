package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var config_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Display merged kubeconfig settings or a specified kubeconfig file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_viewCmd).Standalone()

	config_viewCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	config_viewCmd.Flags().Bool("flatten", false, "Flatten the resulting kubeconfig file into self-contained output (useful for creating portable kubeconfig files)")
	config_viewCmd.Flags().String("merge", "", "Merge the full hierarchy of kubeconfig files")
	config_viewCmd.Flags().Bool("minify", false, "Remove all information not used by current-context from the output")
	config_viewCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	config_viewCmd.Flags().Bool("raw", false, "Display raw byte data and sensitive data")
	config_viewCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	config_viewCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	config_viewCmd.Flag("merge").NoOptDefVal = " "
	configCmd.AddCommand(config_viewCmd)

	carapace.Gen(config_viewCmd).FlagCompletion(carapace.ActionMap{
		"output":   kubectl.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
	})
}
