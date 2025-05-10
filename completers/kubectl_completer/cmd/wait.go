package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var waitCmd = &cobra.Command{
	Use:     "wait ([-f FILENAME] | resource.group/resource.name | resource.group [(-l label | --all)]) [--for=delete|--for condition=available|--for=jsonpath='{}'=value]",
	Short:   "Experimental: Wait for a specific condition on one or many resources",
	GroupID: "advanced",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitCmd).Standalone()

	waitCmd.Flags().Bool("all", false, "Select all resources in the namespace of the specified resource types")
	waitCmd.Flags().BoolP("all-namespaces", "A", false, "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.")
	waitCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	waitCmd.Flags().String("field-selector", "", "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type.")
	waitCmd.Flags().StringSliceP("filename", "f", nil, "identifying the resource.")
	waitCmd.Flags().String("for", "", "The condition to wait on: [create|delete|condition=condition-name[=condition-value]|jsonpath='{JSONPath expression}'=[JSONPath value]]. The default condition-value is true. Condition values are compared after Unicode simple case folding, which is a more general form of case-insensitivity.")
	waitCmd.Flags().Bool("local", false, "If true, annotation will NOT contact api-server but run locally.")
	waitCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	waitCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	waitCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', and '!='.(e.g. -l key1=value1,key2=value2)")
	waitCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	waitCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	waitCmd.Flags().String("timeout", "", "The length of time to wait before giving up. Zero means check once and don't wait, negative means wait for a week.")
	rootCmd.AddCommand(waitCmd)

	carapace.Gen(waitCmd).FlagCompletion(carapace.ActionMap{
		"filename": carapace.ActionFiles(),
		"output":   kubectl.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
	})
}
