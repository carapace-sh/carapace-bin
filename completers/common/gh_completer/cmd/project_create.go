package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var project_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_createCmd).Standalone()

	project_createCmd.Flags().String("format", "", "Output format: {json}")
	project_createCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	project_createCmd.Flags().String("owner", "", "Login of the owner. Use \"@me\" for the current user.")
	project_createCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	project_createCmd.Flags().String("title", "", "Title for the project")
	project_createCmd.MarkFlagRequired("title")
	projectCmd.AddCommand(project_createCmd)

	carapace.Gen(project_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json"),
		"owner":  gh.ActionOwners(gh.HostOpts{}),
	})
}
