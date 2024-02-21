package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/helm"
	"github.com/spf13/cobra"
)

var get_manifestCmd = &cobra.Command{
	Use:   "manifest",
	Short: "download the manifest for a named release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_manifestCmd).Standalone()
	get_manifestCmd.Flags().Int("revision", 0, "get the named release with revision")
	getCmd.AddCommand(get_manifestCmd)

	carapace.Gen(get_manifestCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return helm.ActionReleases(helm.ReleasesOpts{
				Namespace:   rootCmd.Flag("namespace").Value.String(),
				KubeContext: rootCmd.Flag("kube-context").Value.String(),
			})
		}),
	)
}
