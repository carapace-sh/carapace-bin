package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_editCmd = &cobra.Command{
	Use:   "edit {<number> | <url>}",
	Short: "Edit an issue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	issue_editCmd.Flags().StringSlice("add-assignee", nil, "Add assigned users by their `login`. Use \"@me\" to assign yourself.")
	issue_editCmd.Flags().StringSlice("add-label", nil, "Add labels by `name`")
	issue_editCmd.Flags().StringSlice("add-project", nil, "Add the issue to projects by `name`")
	issue_editCmd.Flags().StringP("body", "b", "", "Set the new body.")
	issue_editCmd.Flags().StringP("body-file", "F", "", "Read body text from `file`")
	issue_editCmd.Flags().StringP("milestone", "m", "", "Edit the milestone the issue belongs to by `name`")
	issue_editCmd.Flags().StringSlice("remove-assignee", nil, "Remove assigned users by their `login`. Use \"@me\" to unassign yourself.")
	issue_editCmd.Flags().StringSlice("remove-label", nil, "Remove labels by `name`")
	issue_editCmd.Flags().StringSlice("remove-project", nil, "Remove the issue from projects by `name`")
	issue_editCmd.Flags().StringP("title", "t", "", "Set the new title.")
	issueCmd.AddCommand(issue_editCmd)

	carapace.Gen(issue_editCmd).FlagCompletion(carapace.ActionMap{
		"add-assignee": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionAssignableUsers(issue_editCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"add-label": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionLabels(issue_editCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"add-project": action.ActionProjects(issue_editCmd, action.ProjectOpts{Open: true}),
		"body-file":   carapace.ActionFiles(),
		"milestone":   action.ActionMilestones(issue_editCmd),
		// TODO remove-assignee, remove-label, remove-project
	})

	carapace.Gen(issue_editCmd).PositionalCompletion(
		action.ActionIssues(issue_editCmd, action.IssueOpts{Open: true}),
	)
}
