package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/kubectl"
	"github.com/spf13/cobra"
)

var rollout_statusCmd = &cobra.Command{
	Use:   "status (TYPE NAME | TYPE/NAME) [flags]",
	Short: "Show the status of the rollout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rollout_statusCmd).Standalone()

	rollout_statusCmd.Flags().StringSliceP("filename", "f", nil, "Filename, directory, or URL to files identifying the resource to get from a server.")
	rollout_statusCmd.Flags().StringP("kustomize", "k", "", "Process the kustomization directory. This flag can't be used together with -f or -R.")
	rollout_statusCmd.Flags().BoolP("recursive", "R", false, "Process the directory used in -f, --filename recursively. Useful when you want to manage related manifests organized within the same directory.")
	rollout_statusCmd.Flags().String("revision", "", "Pin to a specific revision for showing its status. Defaults to 0 (last revision).")
	rollout_statusCmd.Flags().StringP("selector", "l", "", "Selector (label query) to filter on, supports '=', '==', '!=', 'in', 'notin'.(e.g. -l key1=value1,key2=value2,key3 in (value3)). Matching objects must satisfy all of the specified label constraints.")
	rollout_statusCmd.Flags().String("timeout", "", "The length of time to wait before ending watch, zero means never. Any other values should contain a corresponding time unit (e.g. 1s, 2m, 3h).")
	rollout_statusCmd.Flags().BoolP("watch", "w", false, "Watch the status of the rollout until it's done.")
	rolloutCmd.AddCommand(rollout_statusCmd)

	carapace.Gen(rollout_statusCmd).FlagCompletion(carapace.ActionMap{
		"filename":  carapace.ActionFiles(),
		"kustomize": carapace.ActionDirectories(),
	})

	carapace.Gen(rollout_statusCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return kubectl.ActionApiResourceResources(kubectl.ApiResourceResourcesOpts{
				Namespace: rootCmd.Flag("namespace").Value.String(),
			})
		}),
	)
}
