package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var eventsCmd = &cobra.Command{
	Use:   "events [(-o|--output=)json|yaml|kyaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-as-json|jsonpath-file] [--for TYPE/NAME] [--watch] [--types=Normal,Warning]",
	Short: "List events",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eventsCmd).Standalone()

	eventsCmd.Flags().BoolP("all-namespaces", "A", false, "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.")
	eventsCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	eventsCmd.Flags().String("chunk-size", "", "Return large lists in chunks rather than all at once. Pass 0 to disable.")
	eventsCmd.Flags().String("for", "", "Filter events to only those pertaining to the specified resource.")
	eventsCmd.Flags().Bool("no-headers", false, "When using the default output format, don't print headers.")
	eventsCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, kyaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	eventsCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	eventsCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	eventsCmd.Flags().StringSlice("types", nil, "Output only events of given types.")
	eventsCmd.Flags().BoolP("watch", "w", false, "After listing the requested events, watch for more events.")
	rootCmd.AddCommand(eventsCmd)

	carapace.Gen(eventsCmd).FlagCompletion(carapace.ActionMap{
		"output": kubectl.ActionOutputFormats(),
		"template": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if eventsCmd.Flag("output").Value.String() == "go-template-file" {
				return carapace.ActionFiles()
			}
			return carapace.ActionValues()
		}),
	})
}
