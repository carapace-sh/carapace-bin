package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var project_itemAddCmd = &cobra.Command{
	Use:   "item-add [<number>]",
	Short: "Add a pull request or an issue to a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_itemAddCmd).Standalone()

	project_itemAddCmd.Flags().String("format", "", "Output format: {json}")
	project_itemAddCmd.Flags().String("owner", "", "Login of the owner. Use \"@me\" for the current user.")
	project_itemAddCmd.Flags().String("url", "", "URL of the issue or pull request to add to the project")
	project_itemAddCmd.MarkFlagRequired("url")
	projectCmd.AddCommand(project_itemAddCmd)

	carapace.Gen(project_itemAddCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json"),
		"owner":  gh.ActionOwners(gh.HostOpts{}),
	})

	carapace.Gen(project_itemAddCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return gh.ActionProjects(gh.ProjectOpts{
				Owner: project_itemAddCmd.Flag("owner").Value.String(),
				Open:  true,
			})
		}),
	)
}
