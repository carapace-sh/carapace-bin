package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var project_fieldCreateCmd = &cobra.Command{
	Use:   "field-create [<number>]",
	Short: "Create a field in a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_fieldCreateCmd).Standalone()

	project_fieldCreateCmd.Flags().String("data-type", "", "DataType of the new field.: {TEXT|SINGLE_SELECT|DATE|NUMBER}")
	project_fieldCreateCmd.Flags().String("format", "", "Output format: {json}")
	project_fieldCreateCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	project_fieldCreateCmd.Flags().String("name", "", "Name of the new field")
	project_fieldCreateCmd.Flags().String("owner", "", "Login of the owner. Use \"@me\" for the current user.")
	project_fieldCreateCmd.Flags().StringSlice("single-select-options", nil, "Options for SINGLE_SELECT data type")
	project_fieldCreateCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	project_fieldCreateCmd.MarkFlagRequired("data-type")
	project_fieldCreateCmd.MarkFlagRequired("name")
	projectCmd.AddCommand(project_fieldCreateCmd)

	carapace.Gen(project_fieldCreateCmd).FlagCompletion(carapace.ActionMap{
		"data-type": carapace.ActionValues("TEXT", "SINGLE_SELECT", "DATE", "NUMBER"),
		"format":    carapace.ActionValues("json"),
		"owner":     gh.ActionOwners(gh.HostOpts{}),
	})

	carapace.Gen(project_fieldCreateCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return gh.ActionProjects(gh.ProjectOpts{
				Owner: project_fieldCreateCmd.Flag("owner").Value.String(),
				Open:  true,
			})
		}),
	)
}
