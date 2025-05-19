package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
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
	project_itemAddCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	project_itemAddCmd.Flags().String("owner", "", "Login of the owner. Use \"@me\" for the current user.")
	project_itemAddCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
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
