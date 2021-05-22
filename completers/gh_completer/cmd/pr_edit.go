package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_editCmd = &cobra.Command{
	Use:   "edit [<number> | <url> | <branch>]",
	Short: "Edit a pull request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	pr_editCmd.Flags().StringSlice("add-assignee", nil, "Add assigned users by their `login`. Use \"@me\" to assign yourself.")
	pr_editCmd.Flags().StringSlice("add-label", nil, "Add labels by `name`")
	pr_editCmd.Flags().StringSlice("add-project", nil, "Add the pull request to projects by `name`")
	pr_editCmd.Flags().StringSlice("add-reviewer", nil, "Add reviewers by their `login`.")
	pr_editCmd.Flags().StringP("base", "B", "", "Change the base `branch` for this pull request")
	pr_editCmd.Flags().StringP("body", "b", "", "Set the new body.")
	pr_editCmd.Flags().StringP("body-file", "F", "", "Read body text from `file`")
	pr_editCmd.Flags().StringP("milestone", "m", "", "Edit the milestone the pull request belongs to by `name`")
	pr_editCmd.Flags().StringSlice("remove-assignee", nil, "Remove assigned users by their `login`. Use \"@me\" to unassign yourself.")
	pr_editCmd.Flags().StringSlice("remove-label", nil, "Remove labels by `name`")
	pr_editCmd.Flags().StringSlice("remove-project", nil, "Remove the pull request from projects by `name`")
	pr_editCmd.Flags().StringSlice("remove-reviewer", nil, "Remove reviewers by their `login`.")
	pr_editCmd.Flags().StringP("title", "t", "", "Set the new title.")
	prCmd.AddCommand(pr_editCmd)

	carapace.Gen(pr_editCmd).FlagCompletion(carapace.ActionMap{
		"add-assignee": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionAssignableUsers(pr_editCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"add-label": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionLabels(pr_editCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"add-project": action.ActionProjects(pr_editCmd, action.ProjectOpts{Open: true}),
		"add-reviewer": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionAssignableUsers(pr_editCmd).Invoke(c).Filter(c.Parts).ToA()
		}),
		"body-file": carapace.ActionFiles(),
		"milestone": action.ActionMilestones(pr_editCmd),
		// TODO remove-reviewer, remove-assignee, remove-label, remove-project
	})

	carapace.Gen(pr_editCmd).PositionalCompletion(
		action.ActionPullRequests(pr_editCmd, action.PullRequestOpts{Open: true}),
	)
}
