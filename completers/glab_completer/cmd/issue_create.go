package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_createCmd = &cobra.Command{
	Use:     "create [flags]",
	Short:   "Create an issue.",
	Aliases: []string{"new"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_createCmd).Standalone()

	issue_createCmd.Flags().StringSliceP("assignee", "a", []string{}, "Assign issue to people by their `usernames`.")
	issue_createCmd.Flags().BoolP("confidential", "c", false, "Set an issue to be confidential. (default false)")
	issue_createCmd.Flags().StringP("description", "d", "", "Issue description.")
	issue_createCmd.Flags().String("due-date", "", "A date in 'YYYY-MM-DD' format.")
	issue_createCmd.Flags().String("epic", "", "ID of the epic to add the issue to.")
	issue_createCmd.Flags().StringSliceP("label", "l", []string{}, "Add label by name. Multiple labels should be comma-separated.")
	issue_createCmd.Flags().String("link-type", "", "Type for the issue link")
	issue_createCmd.Flags().StringSlice("linked-issues", []string{}, "The IIDs of issues that this issue links to.")
	issue_createCmd.Flags().String("linked-mr", "", "The IID of a merge request in which to resolve all issues.")
	issue_createCmd.Flags().StringP("milestone", "m", "", "The global ID or title of a milestone to assign.")
	issue_createCmd.Flags().Bool("no-editor", false, "Don't open editor to enter a description. If set to true, uses prompt. (default false)")
	issue_createCmd.Flags().Bool("recover", false, "Save the options to a file if the issue fails to be created. If the file exists, the options will be loaded from the recovery file. (EXPERIMENTAL.)")
	issue_createCmd.Flags().StringP("time-estimate", "e", "", "Set time estimate for the issue.")
	issue_createCmd.Flags().StringP("time-spent", "s", "", "Set time spent for the issue.")
	issue_createCmd.Flags().StringP("title", "t", "", "Issue title.")
	issue_createCmd.Flags().Bool("web", false, "Continue issue creation with web interface.")
	issue_createCmd.Flags().StringP("weight", "w", "", "Issue weight. Valid values are greater than or equal to 0.")
	issue_createCmd.Flags().BoolP("yes", "y", false, "Don't prompt for confirmation to submit the issue.")
	issueCmd.AddCommand(issue_createCmd)

	carapace.Gen(issue_createCmd).FlagCompletion(carapace.ActionMap{
		// TODO more flags
		"assignee":      action.ActionProjectMembers(issue_createCmd).UniqueList(","),
		"label":         action.ActionLabels(issue_createCmd).UniqueList(","),
		"link-type":     carapace.ActionValues("relates_to", "blocks", "is_blocked_by"),
		"linked-issues": action.ActionIssues(issue_createCmd, "opened").UniqueList(","),
		"linked-mr":     action.ActionMergeRequests(issue_createCmd, "opened"),
		"milestone":     action.ActionMilestones(issue_createCmd),
	})
}
