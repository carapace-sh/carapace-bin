package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var describeCmd = &cobra.Command{
	Use:     "describe (-f FILENAME | TYPE [NAME_PREFIX | -l label] | TYPE/NAME)",
	Short:   "Show details of a specific resource or group of resources",
	GroupID: "troubleshooting",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(describeCmd).Standalone()

	describeCmd.Flags().BoolP("all-namespaces", "A", false, "If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace.")
	describeCmd.Flags().String("chunk-size", "", "Return large lists in chunks rather than all at once. Pass 0 to disable.")
	describeCmd.Flags().StringSliceP("filename", "f", nil, "Filename, directory, or URL to files containing the resource to describe")
	describeCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	describeCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	describeCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.")
	describeCmd.Flags().Bool("show-events", false, "If true, display events related to the described object.")
	rootCmd.AddCommand(describeCmd)

	carapace.Gen(describeCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
	})

	carapace.Gen(describeCmd).PositionalCompletion(
		kubectl.ActionApiResources(),
	)

	carapace.Gen(describeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return kubectl.ActionResources(kubectl.ResourceOpts{
				Context:   rootCmd.Flag("context").Value.String(),
				Namespace: rootCmd.Flag("namespace").Value.String(),
				Types:     c.Args[0],
			}).Filter(c.Args[1:]...)
		}),
	)
}
