package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_editCmd = &cobra.Command{
	Use:     "edit [<number> | <url> | <branch>]",
	Short:   "Edit a pull request",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_editCmd).Standalone()

	pr_editCmd.Flags().StringSlice("add-assignee", nil, "Add assigned users by their `login`. Use \"@me\" to assign yourself, or \"@copilot\" to assign Copilot.")
	pr_editCmd.Flags().StringSlice("add-label", nil, "Add labels by `name`")
	pr_editCmd.Flags().StringSlice("add-project", nil, "Add the pull request to projects by `title`")
	pr_editCmd.Flags().StringSlice("add-reviewer", nil, "Add reviewers by their `login`.")
	pr_editCmd.Flags().StringP("base", "B", "", "Change the base `branch` for this pull request")
	pr_editCmd.Flags().StringP("body", "b", "", "Set the new body.")
	pr_editCmd.Flags().StringP("body-file", "F", "", "Read body text from `file` (use \"-\" to read from standard input)")
	pr_editCmd.Flags().StringP("milestone", "m", "", "Edit the milestone the pull request belongs to by `name`")
	pr_editCmd.Flags().StringSlice("remove-assignee", nil, "Remove assigned users by their `login`. Use \"@me\" to unassign yourself, or \"@copilot\" to unassign Copilot.")
	pr_editCmd.Flags().StringSlice("remove-label", nil, "Remove labels by `name`")
	pr_editCmd.Flags().Bool("remove-milestone", false, "Remove the milestone association from the pull request")
	pr_editCmd.Flags().StringSlice("remove-project", nil, "Remove the pull request from projects by `title`")
	pr_editCmd.Flags().StringSlice("remove-reviewer", nil, "Remove reviewers by their `login`.")
	pr_editCmd.Flags().StringP("title", "t", "", "Set the new title.")
	prCmd.AddCommand(pr_editCmd)

	carapace.Gen(pr_editCmd).FlagCompletion(carapace.ActionMap{
		"add-assignee": action.ActionAssignableUsers(pr_editCmd).UniqueList(","),
		"add-label":    action.ActionLabels(pr_editCmd).UniqueList(","),
		"add-project":  action.ActionProjects(pr_editCmd, action.ProjectOpts{Open: true}),
		"add-reviewer": action.ActionAssignableUsers(pr_editCmd).UniqueList(","),
		"base":         action.ActionBranches(pr_editCmd),
		"body":         action.ActionBody(pr_editCmd),
		"body-file":    carapace.ActionFiles(),
		"milestone":    action.ActionMilestones(pr_editCmd),
		"remove-assignee": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionPullRequestAssignees(pr_editCmd, c.Args[0])
			}
			return carapace.ActionValues()
		}),
		"remove-label": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionPullRequestLabels(pr_editCmd, c.Args[0])
			}
			return carapace.ActionValues()
		}),
		"remove-reviewer": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionPullRequestReviewers(pr_editCmd, c.Args[0])
			}
			return carapace.ActionValues()
		}),
		// TODO remove-project
	})

	carapace.Gen(pr_editCmd).PositionalCompletion(
		action.ActionPullRequests(pr_editCmd, action.PullRequestOpts{Open: true}),
	)
}
