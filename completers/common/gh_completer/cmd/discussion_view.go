package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var discussion_viewCmd = &cobra.Command{
	Use:     "view {<number> | <discussion-url> | <comment-id> | <comment-url>} [flags]",
	Short:   "View a discussion (preview)",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discussion_viewCmd).Standalone()

	discussion_viewCmd.Flags().String("after", "", "Cursor for the next page")
	discussion_viewCmd.Flags().BoolP("comments", "c", false, "View discussion comments")
	discussion_viewCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	discussion_viewCmd.Flags().StringSlice("json", nil, "Output JSON with the specified `fields`")
	discussion_viewCmd.Flags().StringP("limit", "L", "", "Maximum number of comments or replies to fetch")
	discussion_viewCmd.Flags().String("order", "", "Order of comments or replies: {oldest|newest}")
	discussion_viewCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	discussion_viewCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	discussion_viewCmd.Flags().BoolP("web", "w", false, "Open a discussion in the browser")
	discussionCmd.AddCommand(discussion_viewCmd)

	carapace.Gen(discussion_viewCmd).FlagCompletion(carapace.ActionMap{
		"json":  gh.ActionDiscussionViewFields().UniqueList(","),
		"order": carapace.ActionValues("oldest", "newest"),
		"repo":  gh.ActionHostOwnerRepositories(),
	})

	carapace.Gen(discussion_viewCmd).PositionalAnyCompletion(
		action.ActionDiscussions(discussion_viewCmd),
	)
}
