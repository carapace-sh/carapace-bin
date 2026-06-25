package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_editCmd = &cobra.Command{
	Use:     "edit {<numbers> | <urls>}",
	Short:   "Edit issues",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_editCmd).Standalone()

	issue_editCmd.Flags().StringSlice("add-assignee", nil, "Add assigned users by their `login`. Use \"@me\" to assign yourself, or \"@copilot\" to assign Copilot.")
	issue_editCmd.Flags().StringSlice("add-blocked-by", nil, "Add 'blocked by' relationships by issue `number` or URL")
	issue_editCmd.Flags().StringSlice("add-blocking", nil, "Add 'blocking' relationships by issue `number` or URL")
	issue_editCmd.Flags().StringSlice("add-label", nil, "Add labels by `name`")
	issue_editCmd.Flags().StringSlice("add-project", nil, "Add the issue to projects by `title`")
	issue_editCmd.Flags().StringSlice("add-sub-issue", nil, "Add sub-issues by `number` or URL")
	issue_editCmd.Flags().StringP("body", "b", "", "Set the new body.")
	issue_editCmd.Flags().StringP("body-file", "F", "", "Read body text from `file` (use \"-\" to read from standard input)")
	issue_editCmd.Flags().StringP("milestone", "m", "", "Edit the milestone the issue belongs to by `name`")
	issue_editCmd.Flags().String("parent", "", "Set the parent issue by `number` or URL")
	issue_editCmd.Flags().StringSlice("remove-assignee", nil, "Remove assigned users by their `login`. Use \"@me\" to unassign yourself, or \"@copilot\" to unassign Copilot.")
	issue_editCmd.Flags().StringSlice("remove-blocked-by", nil, "Remove 'blocked by' relationships by issue `number` or URL")
	issue_editCmd.Flags().StringSlice("remove-blocking", nil, "Remove 'blocking' relationships by issue `number` or URL")
	issue_editCmd.Flags().StringSlice("remove-label", nil, "Remove labels by `name`")
	issue_editCmd.Flags().Bool("remove-milestone", false, "Remove the milestone association from the issue")
	issue_editCmd.Flags().Bool("remove-parent", false, "Remove the parent issue")
	issue_editCmd.Flags().StringSlice("remove-project", nil, "Remove the issue from projects by `title`")
	issue_editCmd.Flags().StringSlice("remove-sub-issue", nil, "Remove sub-issues by `number` or URL")
	issue_editCmd.Flags().Bool("remove-type", false, "Remove the issue type from the issue")
	issue_editCmd.Flags().StringP("title", "t", "", "Set the new title.")
	issue_editCmd.Flags().String("type", "", "Set the issue type by `name`")
	issueCmd.AddCommand(issue_editCmd)

	carapace.Gen(issue_editCmd).FlagCompletion(carapace.ActionMap{
		"add-assignee":   action.ActionAssignableUsers(issue_editCmd).UniqueList(","),
		"add-blocked-by": action.ActionIssues(issue_editCmd, action.IssueOpts{Open: true}),
		"add-blocking":   action.ActionIssues(issue_editCmd, action.IssueOpts{Open: true}),
		"add-label":      action.ActionLabels(issue_editCmd).UniqueList(","),
		"add-project":    action.ActionProjects(issue_editCmd, action.ProjectOpts{Open: true}),
		"add-sub-issue":  action.ActionIssues(issue_editCmd, action.IssueOpts{Open: true, Closed: true}),
		"body":           action.ActionBody(issue_editCmd),
		"body-file":      carapace.ActionFiles(),
		"milestone":      action.ActionMilestones(issue_editCmd),
		"parent":         action.ActionIssues(issue_editCmd, action.IssueOpts{Open: true, Closed: true}),
		"remove-assignee": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionIssueAssignees(issue_editCmd, c.Args[0])
			}
			return carapace.ActionValues()
		}),
		"remove-blocked-by": action.ActionIssues(issue_editCmd, action.IssueOpts{Open: true, Closed: true}),
		"remove-blocking":   action.ActionIssues(issue_editCmd, action.IssueOpts{Open: true, Closed: true}),
		"remove-label": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionIssueLabels(issue_editCmd, c.Args[0])
			}
			return carapace.ActionValues()
		}),
		"remove-sub-issue": action.ActionIssues(issue_editCmd, action.IssueOpts{Open: true, Closed: true}),
		"type":             action.ActionIssueTypes(issue_editCmd),
		// TODO remove-project
	})

	carapace.Gen(issue_editCmd).PositionalAnyCompletion(
		action.ActionIssues(issue_editCmd, action.IssueOpts{Open: true, Closed: true}),
	)
}
