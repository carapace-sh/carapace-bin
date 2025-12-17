package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var rollout_pauseCmd = &cobra.Command{
	Use:   "pause RESOURCE",
	Short: "Mark the provided resource as paused",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rollout_pauseCmd).Standalone()

	rollout_pauseCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	rollout_pauseCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	rollout_pauseCmd.Flags().StringSliceP("filename", "f", nil, "Filename, directory, or URL to files identifying the resource to get from a server.")
	rollout_pauseCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	rollout_pauseCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, kyaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	rollout_pauseCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	rollout_pauseCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.")
	rollout_pauseCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	rollout_pauseCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	rolloutCmd.AddCommand(rollout_pauseCmd)

	carapace.Gen(rollout_pauseCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    kubectl.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(rollout_pauseCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return kubectl.ActionApiResourceResources(kubectl.ApiResourceResourcesOpts{
				Namespace: rootCmd.Flag("namespace").Value.String(),
			})
		}),
	)
}
