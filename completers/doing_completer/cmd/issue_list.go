package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/doing"
	"github.com/spf13/cobra"
)

var issue_listCmd = &cobra.Command{
	Use:   "list [OPTIONS]",
	Short: "List issues related to the project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_listCmd).Standalone()

	issue_listCmd.Flags().StringP("assignee", "a", "", "Filter by assignee")
	issue_listCmd.Flags().StringP("author", "A", "", "Filter by author")
	issue_listCmd.Flags().Bool("help", false, "Show this message and exit")
	issue_listCmd.Flags().StringP("label", "l", "", "Filter by labels")
	issue_listCmd.Flags().Bool("no-show_state", false, "Do not show column with work item state")
	issue_listCmd.Flags().Bool("no-web", false, "Open overview of issues in the web browser")
	issue_listCmd.Flags().StringP("output_format", "o", "", "Output format")
	issue_listCmd.Flags().Bool("show_state", false, "Show column with work item state")
	issue_listCmd.Flags().StringP("state", "s", "", "Filter by state")
	issue_listCmd.Flags().String("story_points", "", "Filter on number of story points")
	issue_listCmd.Flags().StringP("type", "t", "", "Type of work item")
	issue_listCmd.Flags().BoolP("web", "w", false, "Open overview of issues in the web browser")
	issueCmd.AddCommand(issue_listCmd)

	// TODO completion
	carapace.Gen(issue_listCmd).FlagCompletion(carapace.ActionMap{
		"assignee":      carapace.ActionValues(),
		"author":        carapace.ActionValues(),
		"label":         carapace.ActionValues(),
		"output_format": carapace.ActionValues("table", "array"),
		"state":         carapace.ActionValues(),
		"story_points":  carapace.ActionValues(),
		"type":          doing.ActionWorkItemTypes(),
	})

}
