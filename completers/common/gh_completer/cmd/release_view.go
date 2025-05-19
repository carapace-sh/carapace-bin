package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_viewCmd = &cobra.Command{
	Use:     "view [<tag>]",
	Short:   "View information about a release",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_viewCmd).Standalone()

	release_viewCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	release_viewCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	release_viewCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	release_viewCmd.Flags().BoolP("web", "w", false, "Open the release in the browser")
	releaseCmd.AddCommand(release_viewCmd)

	carapace.Gen(release_viewCmd).FlagCompletion(carapace.ActionMap{
		"json": action.ActionReleaseFields().UniqueList(","),
	})

	carapace.Gen(release_viewCmd).PositionalCompletion(
		action.ActionReleases(release_viewCmd),
	)
}
