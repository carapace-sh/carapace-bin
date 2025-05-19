package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pr_listCmd = &cobra.Command{
	Use:   "list [OPTIONS]",
	Short: "List pull requests related to the project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_listCmd).Standalone()

	pr_listCmd.Flags().StringP("assignee", "a", "", "Filter by assigned reviewers")
	pr_listCmd.Flags().Bool("help", false, "Show this message and exit")
	pr_listCmd.Flags().StringP("label", "l", "", "Filter by labels")
	pr_listCmd.Flags().StringP("limit", "L", "", "Maximum number of items to fetch")
	pr_listCmd.Flags().Bool("no-web", false, "Open overview of issues in the web browser")
	pr_listCmd.Flags().StringP("state", "s", "", "Filter by state")
	pr_listCmd.Flags().BoolP("web", "w", false, "Open overview of issues in the web browser")
	prCmd.AddCommand(pr_listCmd)

	// TODO flag completion
	carapace.Gen(pr_listCmd).FlagCompletion(carapace.ActionMap{
		"assignee": carapace.ActionValues(),
		"label":    carapace.ActionValues(),
		"limit":    carapace.ActionValues(),
		"state":    carapace.ActionValues("open", "closed", "merged", "all"),
	})
}
