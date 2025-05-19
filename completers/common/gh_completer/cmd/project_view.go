package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var project_viewCmd = &cobra.Command{
	Use:   "view [<number>]",
	Short: "View a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_viewCmd).Standalone()

	project_viewCmd.Flags().String("format", "", "Output format: {json}")
	project_viewCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	project_viewCmd.Flags().String("owner", "", "Login of the owner. Use \"@me\" for the current user.")
	project_viewCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	project_viewCmd.Flags().BoolP("web", "w", false, "Open a project in the browser")
	projectCmd.AddCommand(project_viewCmd)

	carapace.Gen(project_viewCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json"),
		"owner":  gh.ActionOwners(gh.HostOpts{}),
	})

	carapace.Gen(project_viewCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return gh.ActionProjects(gh.ProjectOpts{
				Owner:  project_viewCmd.Flag("owner").Value.String(),
				Open:   true,
				Closed: true,
			})
		}),
	)
}
