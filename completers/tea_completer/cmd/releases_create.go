package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var releases_createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a release",
	Aliases: []string{"c"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(releases_createCmd).Standalone()

	releases_createCmd.Flags().StringP("asset", "a", "", "Path to file attachment. Can be specified multiple times")
	releases_createCmd.Flags().BoolP("draft", "d", false, "Is a draft")
	releases_createCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	releases_createCmd.Flags().StringP("note", "n", "", "Release notes")
	releases_createCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	releases_createCmd.Flags().BoolP("prerelease", "p", false, "Is a pre-release")
	releases_createCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	releases_createCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	releases_createCmd.Flags().String("tag", "", "Tag name. If the tag does not exist yet, it will be created by Gitea")
	releases_createCmd.Flags().String("target", "", "Target branch name or commit hash. Defaults to the default branch of the repo")
	releases_createCmd.Flags().StringP("title", "t", "", "Release title")
	releasesCmd.AddCommand(releases_createCmd)

	// TODO completion
	carapace.Gen(releases_createCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})
}
