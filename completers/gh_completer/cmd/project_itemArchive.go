package cmd

import (
	"strconv"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var project_itemArchiveCmd = &cobra.Command{
	Use:   "item-archive [<number>]",
	Short: "Archive an item in a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_itemArchiveCmd).Standalone()

	project_itemArchiveCmd.Flags().String("format", "", "Output format: {json}")
	project_itemArchiveCmd.Flags().String("id", "", "ID of the item to archive")
	project_itemArchiveCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	project_itemArchiveCmd.Flags().String("owner", "", "Login of the owner. Use \"@me\" for the current user.")
	project_itemArchiveCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	project_itemArchiveCmd.Flags().Bool("undo", false, "Unarchive an item")
	project_itemArchiveCmd.MarkFlagRequired("id")
	projectCmd.AddCommand(project_itemArchiveCmd)

	carapace.Gen(project_itemArchiveCmd).FlagCompletion(carapace.ActionMap{
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
				Owner:   project_itemArchiveCmd.Flag("owner").Value.String(),
				Project: project,
			})
		}),
		"owner": gh.ActionOwners(gh.HostOpts{}),
	})

	carapace.Gen(project_itemArchiveCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return gh.ActionProjects(gh.ProjectOpts{
				Owner:  project_itemArchiveCmd.Flag("owner").Value.String(),
				Open:   true,
				Closed: true,
			})
		}),
	)
}
