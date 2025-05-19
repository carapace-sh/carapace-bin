package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var project_copyCmd = &cobra.Command{
	Use:   "copy [<number>]",
	Short: "Copy a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_copyCmd).Standalone()

	project_copyCmd.Flags().Bool("drafts", false, "Include draft issues when copying")
	project_copyCmd.Flags().String("format", "", "Output format: {json}")
	project_copyCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	project_copyCmd.Flags().String("source-owner", "", "Login of the source owner. Use \"@me\" for the current user.")
	project_copyCmd.Flags().String("target-owner", "", "Login of the target owner. Use \"@me\" for the current user.")
	project_copyCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	project_copyCmd.Flags().String("title", "", "Title for the new project")
	project_copyCmd.MarkFlagRequired("title")
	projectCmd.AddCommand(project_copyCmd)

	carapace.Gen(project_copyCmd).FlagCompletion(carapace.ActionMap{
		"format":       carapace.ActionValues("json"),
		"source-owner": gh.ActionOwners(gh.HostOpts{}),
		"target-owner": gh.ActionOwners(gh.HostOpts{}),
	})

	carapace.Gen(project_copyCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return gh.ActionProjects(gh.ProjectOpts{
				Owner:  project_copyCmd.Flag("source-owner").Value.String(),
				Open:   true,
				Closed: true,
			})
		}),
	)
}
