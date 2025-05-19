package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var project_unlinkCmd = &cobra.Command{
	Use:   "unlink [<number>]",
	Short: "Unlink a project from a repository or a team",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_unlinkCmd).Standalone()

	project_unlinkCmd.Flags().String("owner", "", "Login of the owner. Use \"@me\" for the current user.")
	project_unlinkCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	project_unlinkCmd.Flags().StringP("team", "T", "", "The team to be unlinked from this project")
	projectCmd.AddCommand(project_unlinkCmd)

	carapace.Gen(project_unlinkCmd).FlagCompletion(carapace.ActionMap{
		"owner": gh.ActionOwners(gh.HostOpts{}),
		"repo":  gh.ActionHostOwnerRepositories(),
		"team":  action.ActionTeams(project_linkCmd),
	})

	carapace.Gen(project_unlinkCmd).PositionalCompletion(
		action.ActionProjects(project_linkCmd, action.ProjectOpts{Open: true}),
	)
}
