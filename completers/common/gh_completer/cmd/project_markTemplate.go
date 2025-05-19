package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var project_markTemplateCmd = &cobra.Command{
	Use:   "mark-template [<number>]",
	Short: "Mark a project as a template",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_markTemplateCmd).Standalone()

	project_markTemplateCmd.Flags().String("format", "", "Output format: {json}")
	project_markTemplateCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	project_markTemplateCmd.Flags().String("owner", "", "Login of the org owner.")
	project_markTemplateCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	project_markTemplateCmd.Flags().Bool("undo", false, "Unmark the project as a template.")
	projectCmd.AddCommand(project_markTemplateCmd)

	carapace.Gen(project_markTemplateCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json"),
		"owner":  gh.ActionOwners(gh.HostOpts{}),
	})

	carapace.Gen(project_markTemplateCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return gh.ActionProjects(gh.ProjectOpts{
				Owner:  project_markTemplateCmd.Flag("owner").Value.String(),
				Open:   true,
				Closed: true,
			})
		}),
	)
}
