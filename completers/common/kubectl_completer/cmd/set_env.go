package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var set_envCmd = &cobra.Command{
	Use:   "env RESOURCE/NAME KEY_1=VAL_1 ... KEY_N=VAL_N",
	Short: "Update environment variables on a pod template",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_envCmd).Standalone()

	set_envCmd.Flags().Bool("all", false, "If true, select all resources in the namespace of the specified resource types")
	set_envCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	set_envCmd.Flags().StringP("containers", "c", "", "The names of containers in the selected pod templates to change - may use wildcards")
	set_envCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	set_envCmd.Flags().StringSliceP("env", "e", nil, "Specify a key-value pair for an environment variable to set into each container.")
	set_envCmd.Flags().String("field-manager", "", "Name of the manager used to track field ownership.")
	set_envCmd.Flags().StringSliceP("filename", "f", nil, "Filename, directory, or URL to files the resource to update the env")
	set_envCmd.Flags().String("from", "", "The name of a resource from which to inject environment variables")
	set_envCmd.Flags().StringSlice("keys", nil, "Comma-separated list of keys to import from specified resource")
	set_envCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	set_envCmd.Flags().Bool("list", false, "If true, display the environment and any changes in the standard format. this flag will removed when we have kubectl view env.")
	set_envCmd.Flags().Bool("local", false, "If true, set env will NOT contact api-server but run locally.")
	set_envCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, kyaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	set_envCmd.Flags().Bool("overwrite", false, "If true, allow environment to be overwritten, otherwise reject updates that overwrite existing environment.")
	set_envCmd.Flags().String("prefix", "", "Prefix to append to variable names")
	set_envCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	set_envCmd.Flags().Bool("resolve", false, "If true, show secret or configmap references when listing variables")
	set_envCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.")
	set_envCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	set_envCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	set_envCmd.Flag("dry-run").NoOptDefVal = " "
	setCmd.AddCommand(set_envCmd)

	carapace.Gen(set_envCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":  kubectl.ActionDryRunModes(),
		"env":      env.ActionNameValues(false),
		"filename": carapace.ActionFiles(),
		"from":     kubectl.ActionApiResourceResources(kubectl.ApiResourceResourcesOpts{}),
		"output":   kubectl.ActionOutputFormats(),
		"template": carapace.ActionFiles(),
	})

	carapace.Gen(set_envCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return kubectl.ActionApiResourceResources(kubectl.ApiResourceResourcesOpts{
				Namespace: rootCmd.Flag("namespace").Value.String(),
			})
		}),
	)
}
