package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var project_editCmd = &cobra.Command{
	Use:   "edit [<number>]",
	Short: "Edit a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_editCmd).Standalone()

	project_editCmd.Flags().StringP("description", "d", "", "New description of the project")
	project_editCmd.Flags().String("format", "", "Output format: {json}")
	project_editCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	project_editCmd.Flags().String("owner", "", "Login of the owner. Use \"@me\" for the current user.")
	project_editCmd.Flags().String("readme", "", "New readme for the project")
	project_editCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	project_editCmd.Flags().String("title", "", "New title for the project")
	project_editCmd.Flags().String("visibility", "", "Change project visibility: {PUBLIC|PRIVATE}")
	projectCmd.AddCommand(project_editCmd)

	carapace.Gen(project_editCmd).FlagCompletion(carapace.ActionMap{
		"format":     carapace.ActionValues("json"),
		"owner":      gh.ActionOwners(gh.HostOpts{}),
		"visibility": carapace.ActionValues("PUBLIC", "PRIVATE").StyleF(style.ForKeyword),
	})

	carapace.Gen(project_editCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return gh.ActionProjects(gh.ProjectOpts{
				Owner:  project_editCmd.Flag("owner").Value.String(),
				Open:   true,
				Closed: true,
			})
		}),
	)
}
