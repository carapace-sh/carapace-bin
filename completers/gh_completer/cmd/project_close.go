package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var project_closeCmd = &cobra.Command{
	Use:   "close [<number>]",
	Short: "Close a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_closeCmd).Standalone()

	project_closeCmd.Flags().String("format", "", "Output format, must be 'json'")
	project_closeCmd.Flags().String("owner", "", "Login of the owner. Use \"@me\" for the current user.")
	project_closeCmd.Flags().Bool("undo", false, "Reopen a closed project")
	projectCmd.AddCommand(project_closeCmd)

	carapace.Gen(project_closeCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json"),
		"owner":  gh.ActionOwners(gh.HostOpts{}),
	})

	carapace.Gen(project_closeCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return gh.ActionProjects(gh.ProjectOpts{
				Owner: project_closeCmd.Flag("owner").Value.String(),
				Open:  true,
			})
		}),
	)
}
