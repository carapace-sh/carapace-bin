package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_editCmd = &cobra.Command{
	Use:     "edit <tag>",
	Short:   "Edit a release",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_editCmd).Standalone()

	release_editCmd.Flags().String("discussion-category", "", "Start a discussion in the specified category when publishing a draft")
	release_editCmd.Flags().Bool("draft", false, "Save the release as a draft instead of publishing it")
	release_editCmd.Flags().Bool("latest", false, "Explicitly mark the release as \"Latest\"")
	release_editCmd.Flags().StringP("notes", "n", "", "Release notes")
	release_editCmd.Flags().StringP("notes-file", "F", "", "Read release notes from `file` (use \"-\" to read from standard input)")
	release_editCmd.Flags().Bool("prerelease", false, "Mark the release as a prerelease")
	release_editCmd.Flags().String("tag", "", "The name of the tag")
	release_editCmd.Flags().String("target", "", "Target `branch` or full commit SHA (default [main branch])")
	release_editCmd.Flags().StringP("title", "t", "", "Release title")
	release_editCmd.Flags().Bool("verify-tag", false, "Abort in case the git tag doesn't already exist in the remote repository")
	releaseCmd.AddCommand(release_editCmd)

	carapace.Gen(release_editCmd).FlagCompletion(carapace.ActionMap{
		"discussion-category": action.ActionDiscussionCategories(release_editCmd),
		"notes-file":          carapace.ActionFiles(),
		"target":              action.ActionBranches(release_editCmd),
	})

	carapace.Gen(release_editCmd).PositionalCompletion(
		action.ActionReleases(release_editCmd),
	)
}
