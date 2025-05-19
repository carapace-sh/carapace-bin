package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var project_linkCmd = &cobra.Command{
	Use:   "link [<number>]",
	Short: "Link a project to a repository or a team",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_linkCmd).Standalone()

	project_linkCmd.Flags().String("owner", "", "Login of the owner. Use \"@me\" for the current user.")
	project_linkCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	project_linkCmd.Flags().StringP("team", "T", "", "The team to be linked to this project")
	projectCmd.AddCommand(project_linkCmd)

	carapace.Gen(project_linkCmd).FlagCompletion(carapace.ActionMap{
		"owner": gh.ActionOwners(gh.HostOpts{}),
		"repo":  gh.ActionHostOwnerRepositories(),
		"team":  action.ActionTeams(project_linkCmd),
	})

	carapace.Gen(project_linkCmd).PositionalCompletion(
		action.ActionProjects(project_linkCmd, action.ProjectOpts{Open: true}),
	)
}
