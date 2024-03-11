package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var alpha_auth_whoamiCmd = &cobra.Command{
	Use:   "whoami",
	Short: "Experimental: Check self subject attributes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alpha_auth_whoamiCmd).Standalone()

	alpha_auth_whoamiCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	alpha_auth_whoamiCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	alpha_auth_whoamiCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	alpha_auth_whoamiCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	alpha_authCmd.AddCommand(alpha_auth_whoamiCmd)

	carapace.Gen(alpha_auth_whoamiCmd).FlagCompletion(carapace.ActionMap{
		"template": carapace.ActionFiles(),
	})
}
