package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/helm"
	"github.com/spf13/cobra"
)

var get_notesCmd = &cobra.Command{
	Use:   "notes",
	Short: "download the notes for a named release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_notesCmd).Standalone()
	get_notesCmd.Flags().Int("revision", 0, "get the named release with revision")
	getCmd.AddCommand(get_notesCmd)

	carapace.Gen(get_notesCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return helm.ActionReleases(helm.ReleasesOpts{
				Namespace:   rootCmd.Flag("namespace").Value.String(),
				KubeContext: rootCmd.Flag("kube-context").Value.String(),
			})
		}),
	)
}
