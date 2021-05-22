package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_viewCmd = &cobra.Command{
	Use:   "view [<tag>]",
	Short: "View information about a release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	release_viewCmd.Flags().BoolP("web", "w", false, "Open the release in the browser")
	releaseCmd.AddCommand(release_viewCmd)
	// TODO cmdutil.AddJSONFlags(cmd, &opts.Exporter, shared.ReleaseFields)

	carapace.Gen(release_viewCmd).PositionalCompletion(
		action.ActionReleases(release_viewCmd),
	)
}
