package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var set_selectorCmd = &cobra.Command{
	Use:   "selector (-f FILENAME | TYPE NAME) EXPRESSIONS [--resource-version=version]",
	Short: "Set the selector on a resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_selectorCmd).Standalone()

	set_selectorCmd.Flags().Bool("all", false, "Select all resources in the namespace of the specified resource types")
	set_selectorCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	set_selectorCmd.Flags().String("dry-run", "none", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	set_selectorCmd.Flags().String("field-manager", "kubectl-set", "Name of the manager used to track field ownership.")
	set_selectorCmd.Flags().StringSliceP("filename", "f", []string{}, "identifying the resource.")
	set_selectorCmd.Flags().Bool("local", false, "If true, annotation will NOT contact api-server but run locally.")
	set_selectorCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	set_selectorCmd.Flags().BoolP("recursive", "R", true, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	set_selectorCmd.Flags().String("resource-version", "", "If non-empty, the selectors update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource.")
	set_selectorCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	set_selectorCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	set_selectorCmd.Flag("dry-run").NoOptDefVal = "unchanged"
	setCmd.AddCommand(set_selectorCmd)

	carapace.Gen(set_selectorCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  kubectl.ActionDryRunModes(),
		"filename": carapace.ActionFiles(),
		"output":   kubectl.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
	})
}
