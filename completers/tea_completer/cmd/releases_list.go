package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var releases_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List Releases",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(releases_listCmd).Standalone()

	releases_listCmd.Flags().String("limit", "", "specify limit of items per page")
	releases_listCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	releases_listCmd.Flags().StringP("output", "o", "", "Output format. (simple, table, csv, tsv, yaml, json)")
	releases_listCmd.Flags().StringP("page", "p", "", "specify page, default is 1")
	releases_listCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	releases_listCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	releasesCmd.AddCommand(releases_listCmd)

	// TODO completion
	carapace.Gen(releases_listCmd).FlagCompletion(carapace.ActionMap{
		"login":  tea.ActionLogins(),
		"output": tea.ActionOutputFormats(),
		"remote": git.ActionRemotes(),
	})
}
