package cmd

import (
	"strconv"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var project_itemDeleteCmd = &cobra.Command{
	Use:   "item-delete [<number>]",
	Short: "Delete an item from a project by ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_itemDeleteCmd).Standalone()

	project_itemDeleteCmd.Flags().String("format", "", "Output format: {json}")
	project_itemDeleteCmd.Flags().String("id", "", "ID of the item to delete")
	project_itemDeleteCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	project_itemDeleteCmd.Flags().String("owner", "", "Login of the owner. Use \"@me\" for the current user.")
	project_itemDeleteCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	project_itemDeleteCmd.MarkFlagRequired("id")
	projectCmd.AddCommand(project_itemDeleteCmd)

	carapace.Gen(project_itemDeleteCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json"),
		"id": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionValues()
			}

			project, err := strconv.Atoi(c.Args[0])
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}

			return gh.ActionProjectItems(gh.ProjectItemOpts{
				Owner:    project_itemDeleteCmd.Flag("owner").Value.String(),
				Project:  project,
				Archived: false,
			})
		}),
		"owner": gh.ActionOwners(gh.HostOpts{}),
	})

	carapace.Gen(project_itemDeleteCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return gh.ActionProjects(gh.ProjectOpts{
				Owner: project_itemDeleteCmd.Flag("owner").Value.String(),
				Open:  true,
			})
		}),
	)
}
