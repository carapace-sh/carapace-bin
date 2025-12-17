package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get [(-o|--output=)json|yaml|kyaml|name|go-template|go-template-file|template|templatefile|jsonpath|jsonpath-as-json|jsonpath-file|custom-columns|custom-columns-file|wide] (TYPE[.VERSION][.GROUP] [NAME | -l label] | TYPE[.VERSION][.GROUP]/NAME ...) [flags]",
	Short:   "Display one or many resources",
	GroupID: "basic intermediate",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	getCmd.Flags().BoolP("all-namespaces", "A", false, "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.")
	getCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	getCmd.Flags().String("chunk-size", "", "Return large lists in chunks rather than all at once. Pass 0 to disable.")
	getCmd.Flags().String("field-selector", "", "Selector (field query) to filter on, supports '=', '==', and '!='.(e.g. --field-selector key1=value1,key2=value2). The server only supports a limited number of field queries per type.")
	getCmd.Flags().StringSliceP("filename", "f", nil, "Filename, directory, or URL to files identifying the resource to get from a server.")
	getCmd.Flags().Bool("ignore-not-found", false, "If set to true, suppresses NotFound error for specific objects that do not exist. Using this flag with commands that query for collections of resources has no effect when no resources are found.")
	getCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	getCmd.Flags().StringSliceP("label-columns", "L", nil, "Accepts a comma separated list of labels that are going to be presented as columns. Names are case-sensitive. You can also use multiple flag options like -L label1 -L label2...")
	getCmd.Flags().Bool("no-headers", false, "When using the default or custom-column output format, don't print headers (default print headers).")
	getCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, kyaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file, custom-columns, custom-columns-file, wide). See custom columns [https://kubernetes.io/docs/reference/kubectl/#custom-columns], golang template [http://golang.org/pkg/text/template/#pkg-overview] and jsonpath template [https://kubernetes.io/docs/reference/kubectl/jsonpath/].")
	getCmd.Flags().Bool("output-watch-events", false, "Output watch event objects when --watch or --watch-only is used. Existing objects are output as initial ADDED events.")
	getCmd.Flags().String("raw", "", "Raw URI to request from the server.  Uses the transport specified by the kubeconfig file.")
	getCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	getCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.")
	getCmd.Flags().Bool("server-print", false, "If true, have the server return the appropriate table output. Supports extension APIs and CRDs.")
	getCmd.Flags().Bool("show-kind", false, "If present, list the resource type for the requested object(s).")
	getCmd.Flags().Bool("show-labels", false, "When printing, show all labels as the last column (default hide labels column)")
	getCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	getCmd.Flags().String("sort-by", "", "If non-empty, sort list types using this field specification.  The field specification is expressed as a JSONPath expression (e.g. '{.metadata.name}'). The field in the API resource specified by this JSONPath expression must be an integer or a string.")
	getCmd.Flags().String("subresource", "", "If specified, gets the subresource of the requested object.")
	getCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	getCmd.Flags().BoolP("watch", "w", false, "After listing/getting the requested object, watch for changes.")
	getCmd.Flags().Bool("watch-only", false, "Watch for changes to the requested object(s), without listing/getting first.")
	rootCmd.AddCommand(getCmd)

	carapace.Gen(getCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    kubectl.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(getCmd).PositionalCompletion(
		kubectl.ActionApiResources().UniqueList(","),
	)

	carapace.Gen(getCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return kubectl.ActionResources(kubectl.ResourceOpts{
				Context:   rootCmd.Flag("context").Value.String(),
				Namespace: rootCmd.Flag("namespace").Value.String(),
				Types:     c.Args[0],
			}).Filter(c.Args[1:]...)
		}),
	)
}
