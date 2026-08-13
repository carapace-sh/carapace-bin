package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var projects_lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "Show all projects",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(projects_lsCmd).Standalone()

	projects_lsCmd.Flags().String("filter", "", "Filter projects")
	projects_lsCmd.Flags().String("format", "", "Output format")
	projects_lsCmd.Flags().Bool("json", false, "Output as JSON")
	projects_lsCmd.Flags().String("limit", "", "Number of results per page")
	projects_lsCmd.Flags().String("next", "", "Show next page of results")
	projects_lsCmd.Flags().Bool("update-required", false, "Show projects that require updating")

	projectsCmd.AddCommand(projects_lsCmd)
}
