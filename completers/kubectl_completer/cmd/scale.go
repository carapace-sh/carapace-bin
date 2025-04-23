package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var scaleCmd = &cobra.Command{
	Use:     "scale [--resource-version=version] [--current-replicas=count] --replicas=COUNT (-f FILENAME | TYPE NAME)",
	Short:   "Set a new size for a deployment, replica set, or replication controller",
	GroupID: "deploy",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scaleCmd).Standalone()

	scaleCmd.Flags().Bool("all", false, "Select all resources in the namespace of the specified resource types")
	scaleCmd.Flags().Bool("allow-missing-template-keys", false, "If true, ignore any errors in templates when a field or map key is missing in the template. Only applies to golang and jsonpath output formats.")
	scaleCmd.Flags().String("current-replicas", "", "Precondition for current size. Requires that the current size of the resource match this value in order to scale. -1 (default) for no condition.")
	scaleCmd.Flags().String("dry-run", "", "Must be \"none\", \"server\", or \"client\". If client strategy, only print the object that would be sent, without sending it. If server strategy, submit server-side request without persisting the resource.")
	scaleCmd.Flags().StringSliceP("filename", "f", []string{}, "Filename, directory, or URL to files identifying the resource to set a new size")
	scaleCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	scaleCmd.Flags().StringP("output", "o", "", "Output format. One of: (json, yaml, name, go-template, go-template-file, template, templatefile, jsonpath, jsonpath-as-json, jsonpath-file).")
	scaleCmd.Flags().Bool("record", false, "Record current kubectl command in the resource annotation. If set to false, do not record the command. If set to true, record the command. If not set, default to updating the existing annotation value only if one already exists.")
	scaleCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	scaleCmd.Flags().String("replicas", "", "The new desired number of replicas. Required.")
	scaleCmd.Flags().String("resource-version", "", "Precondition for resource version. Requires that the current resource version match this value in order to scale.")
	scaleCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.")
	scaleCmd.Flags().Bool("show-managed-fields", false, "If true, keep the managedFields when printing objects in JSON or YAML format.")
	scaleCmd.Flags().String("template", "", "Template string or path to template file to use when -o=go-template, -o=go-template-file. The template format is golang templates [http://golang.org/pkg/text/template/#pkg-overview].")
	scaleCmd.Flags().String("timeout", "", "The length of time to wait before giving up on a scale operation, zero means don't wait. Any other values should contain a corresponding time unit (e.g. 1s, 2m, 3h).")
	scaleCmd.Flag("dry-run").NoOptDefVal = " "
	scaleCmd.Flag("record").Hidden = true
	scaleCmd.MarkFlagRequired("replicas")
	rootCmd.AddCommand(scaleCmd)

	carapace.Gen(scaleCmd).FlagCompletion(carapace.ActionMap{
		"dry-run":   kubectl.ActionDryRunModes(),
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
		"output":    kubectl.ActionOutputFormats(),
		"template":  carapace.ActionFiles(),
	})

	carapace.Gen(scaleCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues("deployments/", "replicasets/", "replicationcontrollers/").NoSpace()
			case 1:
				return kubectl.ActionResources(kubectl.ResourceOpts{
					Context:   rootCmd.Flag("context").Value.String(),
					Namespace: rootCmd.Flag("namespace").Value.String(),
					Types:     c.Parts[0],
				})
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
