package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_updateCmd = &cobra.Command{
	Use:   "update [<id> | <branch>]",
	Short: "Update a merge request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_updateCmd).Standalone()

	mr_updateCmd.Flags().StringSliceP("assignee", "a", nil, "Assign users via username. Prefix with '!' or '-' to remove from existing assignees, '+' to add. Otherwise, replace existing assignees with given users.")
	mr_updateCmd.Flags().StringP("description", "d", "", "Merge request description. Set to \"-\" to open an editor.")
	mr_updateCmd.Flags().Bool("draft", false, "Mark merge request as a draft.")
	mr_updateCmd.Flags().BoolP("fill", "f", false, "Do not prompt for title or body, and just use commit info.")
	mr_updateCmd.Flags().Bool("fill-commit-body", false, "Fill body with each commit body when multiple commits. Can only be used with --fill.")
	mr_updateCmd.Flags().StringSliceP("label", "l", nil, "Add labels.")
	mr_updateCmd.Flags().Bool("lock-discussion", false, "Lock discussion on merge request.")
	mr_updateCmd.Flags().StringP("milestone", "m", "", "Title of the milestone to assign. Set to \"\" or 0 to unassign.")
	mr_updateCmd.Flags().BoolP("ready", "r", false, "Mark merge request as ready to be reviewed and merged.")
	mr_updateCmd.Flags().Bool("remove-source-branch", false, "Toggles the removal of the source branch on merge.")
	mr_updateCmd.Flags().StringSlice("reviewer", nil, "Request review from users by their usernames. Prefix with '!' or '-' to remove from existing reviewers, '+' to add. Otherwise, replace existing reviewers with given users.")
	mr_updateCmd.Flags().Bool("squash-before-merge", false, "Toggles the option to squash commits into a single commit when merging.")
	mr_updateCmd.Flags().String("target-branch", "", "Set target branch.")
	mr_updateCmd.Flags().StringP("title", "t", "", "Title of merge request.")
	mr_updateCmd.Flags().Bool("unassign", false, "Unassign all users.")
	mr_updateCmd.Flags().StringSliceP("unlabel", "u", nil, "Remove labels.")
	mr_updateCmd.Flags().Bool("unlock-discussion", false, "Unlock discussion on merge request.")
	mr_updateCmd.Flags().Bool("wip", false, "Mark merge request as a work in progress. Alternative to --draft.")
	mr_updateCmd.Flags().BoolP("yes", "y", false, "Skip confirmation prompt.")
	mrCmd.AddCommand(mr_updateCmd)

	carapace.Gen(mr_updateCmd).FlagCompletion(carapace.ActionMap{
		"assignee":      action.ActionProjectMembers(mr_updateCmd).UniqueList(","),
		"label":         action.ActionLabels(mr_updateCmd).UniqueList(","),
		"milestone":     action.ActionMilestones(mr_updateCmd),
		"reviewer":      action.ActionProjectMembers(mr_updateCmd).UniqueList(","),
		"target-branch": action.ActionBranches(mr_updateCmd),
		"unlabel":       action.ActionLabels(mr_updateCmd).UniqueList(","),
	})

	carapace.Gen(mr_updateCmd).PositionalCompletion(
		action.ActionMergeRequestsAndBranches(mr_updateCmd, ""),
	)
}
