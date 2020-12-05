package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/kubectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "display one or many resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	getCmd.Flags().BoolP("all-namespaces", "A", false, "list the requested objects across all namespaces")
	getCmd.Flags().Bool("allow-missing-template-keys", false, "ignore any errors in templates when a field or map key is missing in the template")
	getCmd.Flags().String("chunk-size", "", "return large lists in chunks rather than all at once")
	getCmd.Flags().String("field-selector", "", "selector to filter on")
	getCmd.Flags().StringP("filename", "f", "", "file identifying the resource to get from a server")
	getCmd.Flags().Bool("ignore-not-found", false, "f the requested object does not exist exit code 0")
	getCmd.Flags().StringP("kustomize", "k", "", "process the kustomization directory")
	getCmd.Flags().StringP("label-columns", "L", "", "list of labels that are going to be presented as columns")
	getCmd.Flags().Bool("no-headers", false, "don't print headers")
	getCmd.Flags().StringP("output", "o", "", "output format")
	getCmd.Flags().Bool("output-watch-events", false, "output watch event objects")
	getCmd.Flags().String("raw", "", "raw URI to request from the server")
	getCmd.Flags().BoolP("recursive", "R", false, "process the directory used in -f, --filename recursively")
	getCmd.Flags().StringP("selector", "l", "", "Selector to filter on")
	getCmd.Flags().Bool("server-print", false, "return the appropriate table output")
	getCmd.Flags().Bool("show-kind", false, "list the resource type for the requested objects")
	getCmd.Flags().Bool("show-labels", false, "show all labels as the last column")
	getCmd.Flags().String("sort-by", "", "sort list types using this field specification")
	getCmd.Flags().String("template", "", "template string or path to template file")
	getCmd.Flags().BoolP("watch", "w", false, "watch for change")
	getCmd.Flags().Bool("watch-only", false, "watch for changes without getting first")
	rootCmd.AddCommand(getCmd)

	carapace.Gen(getCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(""),
		"kustomize": carapace.ActionDirectories(),
		"output":    carapace.ActionValues("json", "yaml", "wide", "custom-columns=", "custom-columns-file=", "go-template=", "go-template-file=", "jsonpath=", "jsonpath-file="),
		"template":  carapace.ActionFiles(""),
	})

	carapace.Gen(getCmd).PositionalCompletion(
		carapace.ActionMultiParts(",", func(args, parts []string) carapace.Action {
			return action.ActionApiResources().Invoke(args).Filter(parts).ToA()
		}),
	)

	carapace.Gen(getCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(args []string) carapace.Action {
			return action.ActionResources(args[0]).Invoke(args).Filter(args[1:]).ToA()
		}),
	)
}
