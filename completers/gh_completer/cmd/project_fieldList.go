package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var project_fieldListCmd = &cobra.Command{
	Use:   "field-list [<number>]",
	Short: "List the fields in a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_fieldListCmd).Standalone()

	project_fieldListCmd.Flags().String("format", "", "Output format: {json}")
	project_fieldListCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	project_fieldListCmd.Flags().StringP("limit", "L", "", "Maximum number of fields to fetch")
	project_fieldListCmd.Flags().String("owner", "", "Login of the owner. Use \"@me\" for the current user.")
	project_fieldListCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	projectCmd.AddCommand(project_fieldListCmd)

	carapace.Gen(project_fieldListCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json"),
		"owner":  gh.ActionOwners(gh.HostOpts{}),
	})

	carapace.Gen(project_fieldListCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return gh.ActionProjects(gh.ProjectOpts{
				Owner: project_fieldListCmd.Flag("owner").Value.String(),
				Open:  true,
			})
		}),
	)
}
