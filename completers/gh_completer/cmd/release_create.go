package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_createCmd = &cobra.Command{
	Use:     "create [<tag>] [<files>...]",
	Short:   "Create a new release",
	GroupID: "general",
	Aliases: []string{"new"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_createCmd).Standalone()

	release_createCmd.Flags().String("discussion-category", "", "Start a discussion in the specified category")
	release_createCmd.Flags().BoolP("draft", "d", false, "Save the release as a draft instead of publishing it")
	release_createCmd.Flags().Bool("generate-notes", false, "Automatically generate title and notes for the release")
	release_createCmd.Flags().Bool("latest", false, "Mark this release as \"Latest\" (default: automatic based on date and version)")
	release_createCmd.Flags().StringP("notes", "n", "", "Release notes")
	release_createCmd.Flags().StringP("notes-file", "F", "", "Read release notes from `file` (use \"-\" to read from standard input)")
	release_createCmd.Flags().String("notes-start-tag", "", "Tag to use as the starting point for generating release notes")
	release_createCmd.Flags().BoolP("prerelease", "p", false, "Mark the release as a prerelease")
	release_createCmd.Flags().String("target", "", "Target `branch` or full commit SHA (default: main branch)")
	release_createCmd.Flags().StringP("title", "t", "", "Release title")
	release_createCmd.Flags().Bool("verify-tag", false, "Abort in case the git tag doesn't already exist in the remote repository")
	releaseCmd.AddCommand(release_createCmd)

	carapace.Gen(release_createCmd).FlagCompletion(carapace.ActionMap{
		"discussion-category": action.ActionDiscussionCategories(release_createCmd),
		"notes-file":          carapace.ActionFiles(),
		"notes-start-tag":     action.ActionReleases(release_createCmd),
		"target":              action.ActionBranches(release_createCmd),
	})

	carapace.Gen(release_createCmd).PositionalCompletion(
		action.ActionNextReleases(release_createCmd),
	)
	carapace.Gen(release_createCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
