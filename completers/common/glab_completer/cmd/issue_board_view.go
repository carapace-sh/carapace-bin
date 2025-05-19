package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_board_viewCmd = &cobra.Command{
	Use:   "view [flags]",
	Short: "View project issue board.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_board_viewCmd).Standalone()

	issue_board_viewCmd.Flags().StringP("assignee", "a", "", "Filter board issues by assignee username.")
	issue_board_viewCmd.Flags().StringSliceP("labels", "l", nil, "Filter board issues by labels, comma separated.")
	issue_board_viewCmd.Flags().StringP("milestone", "m", "", "Filter board issues by milestone.")
	issue_boardCmd.AddCommand(issue_board_viewCmd)

	carapace.Gen(issue_board_viewCmd).FlagCompletion(carapace.ActionMap{
		"assignee":  action.ActionProjectMembers(issue_board_viewCmd),
		"labels":    action.ActionLabels(issue_board_viewCmd).UniqueList(","),
		"milestone": action.ActionMilestones(issue_board_viewCmd),
	})
}
