package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var releases_editCmd = &cobra.Command{
	Use:     "edit",
	Short:   "Edit one or more releases",
	Aliases: []string{"e"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(releases_editCmd).Standalone()

	releases_editCmd.Flags().StringP("draft", "d", "", "Mark as Draft [True/false]")
	releases_editCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	releases_editCmd.Flags().StringP("note", "n", "", "Change Notes")
	releases_editCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	releases_editCmd.Flags().StringP("prerelease", "p", "", "Mark as Pre-Release [True/false]")
	releases_editCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	releases_editCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	releases_editCmd.Flags().String("tag", "", "Change Tag")
	releases_editCmd.Flags().String("target", "", "Change Target")
	releases_editCmd.Flags().StringP("title", "t", "", "Change Title")
	releasesCmd.AddCommand(releases_editCmd)

	// TODO completion
	carapace.Gen(releases_editCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})
}
