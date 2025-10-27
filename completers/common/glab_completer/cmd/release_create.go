package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_createCmd = &cobra.Command{
	Use:   "create <tag> [<files>...]",
	Short: "Create a new GitLab release, or update an existing one.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_createCmd).Standalone()

	release_createCmd.Flags().StringP("assets-links", "a", "", "JSON string representation of assets links. See documentation for example.")
	release_createCmd.Flags().String("experimental-notes-text-or-file", "", "(EXPERIMENTAL) Value to use as release notes. If a file exists with this value as path, its content will be used. Otherwise, the value itself will be used as text.")
	release_createCmd.Flags().StringSliceP("milestone", "m", nil, "The title of each milestone the release is associated with.")
	release_createCmd.Flags().StringP("name", "n", "", "The release name or title.")
	release_createCmd.Flags().Bool("no-close-milestone", false, "Prevent closing milestones after creating the release.")
	release_createCmd.Flags().Bool("no-update", false, "Prevent updating the existing release.")
	release_createCmd.Flags().StringP("notes", "N", "", "The release notes or description. Accepts Markdown.")
	release_createCmd.Flags().StringP("notes-file", "F", "", "Read release notes 'file'. To read from stdin, use '-'.")
	release_createCmd.Flags().String("package-name", "", "The package name, when uploading assets to the generic package release with --use-package-registry.")
	release_createCmd.Flags().Bool("publish-to-catalog", false, "(EXPERIMENTAL) Publish the release to the GitLab CI/CD catalog.")
	release_createCmd.Flags().StringP("ref", "r", "", "If the specified tag doesn't exist, create a release from the ref and tag it with the specified tag name. Accepts a commit SHA, tag name, or branch name.")
	release_createCmd.Flags().StringP("released-at", "D", "", "ISO 8601 datetime when the release was ready. Defaults to the current datetime.")
	release_createCmd.Flags().StringP("tag-message", "T", "", "Message to use if creating a new annotated tag.")
	release_createCmd.Flags().Bool("use-package-registry", false, "Upload release assets to the generic package registry of the project. Overrides the GITLAB_RELEASE_ASSETS_USE_PACKAGE_REGISTRY environment variable.")
	release_createCmd.Flag("experimental-notes-text-or-file").Hidden = true
	releaseCmd.AddCommand(release_createCmd)

	carapace.Gen(release_createCmd).FlagCompletion(carapace.ActionMap{
		"experimental-notes-text-or-file": carapace.ActionFiles(),
		"milestone":                       action.ActionMilestones(release_createCmd),
		"notes-file":                      carapace.ActionFiles(),
		"ref":                             action.ActionBranches(release_createCmd), // TODO refs
		// TODO released-at
	})

	carapace.Gen(release_createCmd).PositionalCompletion(
		action.ActionReleases(release_createCmd),
	)

	carapace.Gen(release_createCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
