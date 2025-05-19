package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var project_itemListCmd = &cobra.Command{
	Use:   "item-list [<number>]",
	Short: "List the items in a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_itemListCmd).Standalone()

	project_itemListCmd.Flags().String("format", "", "Output format: {json}")
	project_itemListCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	project_itemListCmd.Flags().StringP("limit", "L", "", "Maximum number of items to fetch")
	project_itemListCmd.Flags().String("owner", "", "Login of the owner. Use \"@me\" for the current user.")
	project_itemListCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	projectCmd.AddCommand(project_itemListCmd)

	carapace.Gen(project_itemListCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json"),
		"owner":  gh.ActionOwners(gh.HostOpts{}),
	})

	carapace.Gen(project_itemListCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return gh.ActionProjects(gh.ProjectOpts{
				Owner: project_itemListCmd.Flag("owner").Value.String(),
				Open:  true,
			})
		}),
	)
}
