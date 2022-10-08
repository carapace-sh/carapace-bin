package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var alpha_eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Experimental: List events",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alpha_eventsCmd).Standalone()
	alpha_eventsCmd.Flags().BoolP("all-namespaces", "A", false, "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.")
	alpha_eventsCmd.Flags().Bool("allow-missing-template-keys", true, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	alpha_eventsCmd.Flags().Int64("chunk-size", 500, "Return large lists in chunks rather than all at once. Pass 0 to disable. This flag is beta and may change in the future.")
	alpha_eventsCmd.Flags().String("for", "", "Filter events to only those pertaining to the specified resource.")
	alpha_eventsCmd.Flags().Bool("no-headers", false, "When using the default output format, don't print headers.")
	alpha_eventsCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	alpha_eventsCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	alpha_eventsCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	alpha_eventsCmd.Flags().StringSlice("types", []string{}, "Output only events of given types.")
	alpha_eventsCmd.Flags().BoolP("watch", "w", false, "After listing the requested events, watch for more events.")
	alphaCmd.AddCommand(alpha_eventsCmd)
}
