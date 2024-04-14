package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/doing"
	"github.com/spf13/cobra"
)

var issue_createCmd = &cobra.Command{
	Use:   "create [OPTIONS] ISSUE",
	Short: "Create an issue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_createCmd).Standalone()

	issue_createCmd.Flags().Bool("add-to-current-sprint", false, "If the item needs to be added to the current sprint")
	issue_createCmd.Flags().StringP("assignee", "a", "", "Emailadres or alias of person to assign")
	issue_createCmd.Flags().StringP("body", "b", "", "Optional description of the work item")
	issue_createCmd.Flags().Bool("do-not-add-to-current-sprint", false, "If the item needs to be added to the current sprint")
	issue_createCmd.Flags().Bool("help", false, "Show this message and exit")
	issue_createCmd.Flags().StringP("label", "l", "", "Attach tags (labels) to work item")
	issue_createCmd.Flags().BoolP("mine", "m", false, "Assign issue to yourself")
	issue_createCmd.Flags().Bool("not-mine", false, "Do not assign issue to yourself")
	issue_createCmd.Flags().StringP("parent", "p", "", "To create a child work item, specify the ID of the parent work item")
	issue_createCmd.Flags().StringP("story_points", "s", "", "The number of story points to assign assigned if not specified")
	issue_createCmd.Flags().StringP("type", "t", "", "Type of work item")
	issueCmd.AddCommand(issue_createCmd)

	// TODO completion
	carapace.Gen(issue_createCmd).FlagCompletion(carapace.ActionMap{
		"assignee":     carapace.ActionValues("@me"),
		"body":         carapace.ActionValues(),
		"label":        carapace.ActionValues(),
		"parent":       carapace.ActionValues(),
		"story_points": carapace.ActionValues(),
		"type":         doing.ActionWorkItemTypes(),
	})

	// TODO positional completion
}
