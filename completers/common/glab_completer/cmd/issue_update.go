package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_updateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update issue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_updateCmd).Standalone()

	issue_updateCmd.Flags().StringSliceP("assignee", "a", nil, "Assign users by username. Prefix with '!' or '-' to remove from existing assignees, or '+' to add new. Otherwise, replace existing assignees with these users. Multiple usernames can be comma-separated or specified by repeating the flag.")
	issue_updateCmd.Flags().StringSlice("attach", nil, "(EXPERIMENTAL) Upload a file and reference it at the end of the description. Use \"-\" to read the file from standard input. Repeat the flag to attach multiple files.")
	issue_updateCmd.Flags().BoolP("confidential", "c", false, "Make issue confidential.")
	issue_updateCmd.Flags().StringP("description", "d", "", "Issue description. Set to \"-\" to open an editor.")
	issue_updateCmd.Flags().String("description-file", "", "Read the issue description from a file. Use \"-\" to read from standard input.")
	issue_updateCmd.Flags().String("due-date", "", "A date in 'YYYY-MM-DD' format.")
	issue_updateCmd.Flags().StringSliceP("label", "l", nil, "Add labels.")
	issue_updateCmd.Flags().Bool("lock-discussion", false, "Lock discussion on issue.")
	issue_updateCmd.Flags().StringP("milestone", "m", "", "Title of the milestone to assign Set to \"\" or 0 to unassign.")
	issue_updateCmd.Flags().BoolP("public", "p", false, "Make issue public.")
	issue_updateCmd.Flags().StringP("title", "t", "", "Title of issue.")
	issue_updateCmd.Flags().Bool("unassign", false, "Unassign all users.")
	issue_updateCmd.Flags().StringSliceP("unlabel", "u", nil, "Remove labels.")
	issue_updateCmd.Flags().Bool("unlock-discussion", false, "Unlock discussion on issue.")
	issue_updateCmd.Flags().StringP("weight", "w", "", "Set weight of the issue.")
	issueCmd.AddCommand(issue_updateCmd)

	carapace.Gen(issue_updateCmd).FlagCompletion(carapace.ActionMap{
		"assignee":         action.ActionProjectMembers(issue_updateCmd).UniqueList(","),
		"attach":           carapace.ActionFiles(),
		"description-file": carapace.ActionFiles(),
		"label":            action.ActionLabels(issue_updateCmd).UniqueList(","),
		"milestone":        action.ActionMilestones(issue_updateCmd),
		"unlabel":          action.ActionLabels(issue_updateCmd).UniqueList(","),
	})
	carapace.Gen(issue_updateCmd).PositionalCompletion(
		action.ActionIssues(issue_updateCmd, "opened"),
	)
}
