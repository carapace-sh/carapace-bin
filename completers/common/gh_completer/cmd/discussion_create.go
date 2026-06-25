package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var discussion_createCmd = &cobra.Command{
	Use:     "create [flags]",
	Short:   "Create a new discussion (preview)",
	GroupID: "General commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discussion_createCmd).Standalone()

	discussion_createCmd.Flags().StringP("body", "b", "", "Body for the discussion")
	discussion_createCmd.Flags().StringP("body-file", "F", "", "Read body text from file (use \"-\" to read from stdin)")
	discussion_createCmd.Flags().StringP("category", "c", "", "Category name or slug for the discussion")
	discussion_createCmd.Flags().StringSliceP("label", "l", nil, "Labels to apply to the discussion")
	discussion_createCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	discussion_createCmd.Flags().StringP("title", "t", "", "Title for the discussion")
	discussionCmd.AddCommand(discussion_createCmd)

	carapace.Gen(discussion_createCmd).FlagCompletion(carapace.ActionMap{
		"body-file": carapace.ActionFiles(),
		"category":  action.ActionDiscussionCategories(discussion_createCmd),
		"label":     action.ActionLabels(discussion_createCmd).UniqueList(","),
		"repo":      gh.ActionHostOwnerRepositories(),
	})
}
