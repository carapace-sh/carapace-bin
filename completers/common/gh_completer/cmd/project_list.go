package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var project_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List the projects for an owner",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_listCmd).Standalone()

	project_listCmd.Flags().Bool("closed", false, "Include closed projects")
	project_listCmd.Flags().String("format", "", "Output format: {json}")
	project_listCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	project_listCmd.Flags().StringP("limit", "L", "", "Maximum number of projects to fetch")
	project_listCmd.Flags().String("owner", "", "Login of the owner")
	project_listCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	project_listCmd.Flags().BoolP("web", "w", false, "Open projects list in the browser")
	projectCmd.AddCommand(project_listCmd)

	carapace.Gen(project_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json"),
		"owner":  gh.ActionOwners(gh.HostOpts{}),
	})
}
