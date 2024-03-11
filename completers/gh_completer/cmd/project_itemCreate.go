package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var project_itemCreateCmd = &cobra.Command{
	Use:   "item-create [<number>]",
	Short: "Create a draft issue item in a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_itemCreateCmd).Standalone()

	project_itemCreateCmd.Flags().String("body", "", "Body for the draft issue")
	project_itemCreateCmd.Flags().String("format", "", "Output format: {json}")
	project_itemCreateCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	project_itemCreateCmd.Flags().String("owner", "", "Login of the owner. Use \"@me\" for the current user.")
	project_itemCreateCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	project_itemCreateCmd.Flags().String("title", "", "Title for the draft issue")
	project_itemCreateCmd.MarkFlagRequired("title")
	projectCmd.AddCommand(project_itemCreateCmd)

	carapace.Gen(project_itemCreateCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json"),
		"owner":  gh.ActionOwners(gh.HostOpts{}),
	})

	carapace.Gen(project_itemCreateCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return gh.ActionProjects(gh.ProjectOpts{
				Owner: project_itemCreateCmd.Flag("owner").Value.String(),
				Open:  true,
			})
		}),
	)
}
