package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_viewCmd = &cobra.Command{
	Use:   "view <tag>",
	Short: "View information about a GitLab release.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_viewCmd).Standalone()

	release_viewCmd.Flags().BoolP("web", "w", false, "Open the release in the browser.")
	releaseCmd.AddCommand(release_viewCmd)

	carapace.Gen(release_viewCmd).PositionalCompletion(
		action.ActionReleases(release_viewCmd),
	)
}
