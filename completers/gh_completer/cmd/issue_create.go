package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a new issue",
	GroupID: "General commands",
	Aliases: []string{"new"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_createCmd).Standalone()

	issue_createCmd.Flags().StringSliceP("assignee", "a", []string{}, "Assign people by their `login`. Use \"@me\" to self-assign.")
	issue_createCmd.Flags().StringP("body", "b", "", "Supply a body. Will prompt for one otherwise.")
	issue_createCmd.Flags().StringP("body-file", "F", "", "Read body text from `file` (use \"-\" to read from standard input)")
	issue_createCmd.Flags().BoolP("editor", "e", false, "Skip prompts and open the text editor to write the title and body in. The first line is the title and the remaining text is the body.")
	issue_createCmd.Flags().StringSliceP("label", "l", []string{}, "Add labels by `name`")
	issue_createCmd.Flags().StringP("milestone", "m", "", "Add the issue to a milestone by `name`")
	issue_createCmd.Flags().StringSliceP("project", "p", []string{}, "Add the issue to projects by `title`")
	issue_createCmd.Flags().String("recover", "", "Recover input from a failed run of create")
	issue_createCmd.Flags().StringP("template", "T", "", "Template `name` to use as starting body text")
	issue_createCmd.Flags().StringP("title", "t", "", "Supply a title. Will prompt for one otherwise.")
	issue_createCmd.Flags().BoolP("web", "w", false, "Open the browser to create an issue")
	issueCmd.AddCommand(issue_createCmd)

	carapace.Gen(issue_createCmd).FlagCompletion(carapace.ActionMap{
		"assignee":  action.ActionAssignableUsers(issue_createCmd).UniqueList(","),
		"body":      action.ActionBody(issue_createCmd),
		"body-file": carapace.ActionFiles(),
		"label":     action.ActionLabels(issue_createCmd).UniqueList(","),
		"milestone": action.ActionMilestones(issue_createCmd),
		"project":   action.ActionProjects(issue_createCmd, action.ProjectOpts{Open: true}),
		"template":  action.ActionIssueTemplates(issue_createCmd),
	})
}
