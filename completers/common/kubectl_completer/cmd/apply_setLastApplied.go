package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var apply_setLastAppliedCmd = &cobra.Command{
	Use:   "set-last-applied -f FILENAME",
	Short: "Set the last-applied-configuration annotation on a live object to match the contents of a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apply_setLastAppliedCmd).Standalone()

	apply_setLastAppliedCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	apply_setLastAppliedCmd.Flags().Bool("create-annotation", false, "Will create 'last-applied-configuration' annotations if current objects doesn't have one")
	apply_setLastAppliedCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	apply_setLastAppliedCmd.Flags().StringSliceP("filename", "f", nil, "Filename, directory, or URL to files that contains the last-applied-configuration annotations")
	apply_setLastAppliedCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	apply_setLastAppliedCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	apply_setLastAppliedCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	apply_setLastAppliedCmd.Flag("dry-run").NoOptDefVal = " "
	applyCmd.AddCommand(apply_setLastAppliedCmd)

	carapace.Gen(apply_setLastAppliedCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  kubectl.ActionDryRunModes(),
		"filename": carapace.ActionFiles(),
		"output":   kubectl.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
	})
}
