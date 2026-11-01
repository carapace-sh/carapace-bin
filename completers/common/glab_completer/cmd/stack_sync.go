package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var stack_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Push the stack to GitLab, and create or update its merge requests. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_syncCmd).Standalone()

	stack_syncCmd.Flags().StringSliceP("assignee", "a", nil, "Assign merge request to people by their `usernames`. Multiple usernames can be comma-separated or specified by repeating the flag.")
	stack_syncCmd.Flags().StringSliceP("label", "l", nil, "Add label by `name`. Multiple labels can be comma-separated or specified by repeating the flag.")
	stack_syncCmd.Flags().Bool("no-verify", false, "Bypass the pre-push hook. (See githooks(5) for more information.)")
	stack_syncCmd.Flags().StringSlice("reviewer", nil, "Request review from users by their `usernames`. Multiple usernames can be comma-separated or specified by repeating the flag.")
	stack_syncCmd.Flags().Bool("skip-mr-creation", false, "Skip creating merge requests for branches that don't have one yet.")
	stack_syncCmd.Flags().Bool("skip-push", false, "Rebase the stack locally without pushing branches or creating merge requests. Still fetches from the remote and calls the GitLab API.")
	stack_syncCmd.Flags().Bool("update-base", false, "Rebase the stack onto the latest version of the base branch.")
	stackCmd.AddCommand(stack_syncCmd)

	carapace.Gen(stack_syncCmd).FlagCompletion(carapace.ActionMap{
		"assignee": action.ActionProjectMembers(stack_syncCmd),
		"label":    action.ActionLabels(stack_syncCmd),
		"reviewer": action.ActionProjectMembers(stack_syncCmd),
	})
}
