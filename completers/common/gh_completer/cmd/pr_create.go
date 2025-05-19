package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a pull request",
	GroupID: "General commands",
	Aliases: []string{"new"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_createCmd).Standalone()

	pr_createCmd.Flags().StringSliceP("assignee", "a", nil, "Assign people by their `login`. Use \"@me\" to self-assign.")
	pr_createCmd.Flags().StringP("base", "B", "", "The `branch` into which you want your code merged")
	pr_createCmd.Flags().StringP("body", "b", "", "Body for the pull request")
	pr_createCmd.Flags().StringP("body-file", "F", "", "Read body text from `file` (use \"-\" to read from standard input)")
	pr_createCmd.Flags().BoolP("draft", "d", false, "Mark pull request as a draft")
	pr_createCmd.Flags().Bool("dry-run", false, "Print details instead of creating the PR. May still push git changes.")
	pr_createCmd.Flags().BoolP("editor", "e", false, "Skip prompts and open the text editor to write the title and body in. The first line is the title and the remaining text is the body.")
	pr_createCmd.Flags().BoolP("fill", "f", false, "Use commit info for title and body")
	pr_createCmd.Flags().Bool("fill-first", false, "Use first commit info for title and body")
	pr_createCmd.Flags().Bool("fill-verbose", false, "Use commits msg+body for description")
	pr_createCmd.Flags().StringP("head", "H", "", "The `branch` that contains commits for your pull request (default [current branch])")
	pr_createCmd.Flags().StringSliceP("label", "l", nil, "Add labels by `name`")
	pr_createCmd.Flags().StringP("milestone", "m", "", "Add the pull request to a milestone by `name`")
	pr_createCmd.Flags().Bool("no-maintainer-edit", false, "Disable maintainer's ability to modify pull request")
	pr_createCmd.Flags().StringSliceP("project", "p", nil, "Add the pull request to projects by `title`")
	pr_createCmd.Flags().String("recover", "", "Recover input from a failed run of create")
	pr_createCmd.Flags().StringSliceP("reviewer", "r", nil, "Request reviews from people or teams by their `handle`")
	pr_createCmd.Flags().StringP("template", "T", "", "Template `file` to use as starting body text")
	pr_createCmd.Flags().StringP("title", "t", "", "Title for the pull request")
	pr_createCmd.Flags().BoolP("web", "w", false, "Open the web browser to create a pull request")
	prCmd.AddCommand(pr_createCmd)

	carapace.Gen(pr_createCmd).FlagCompletion(carapace.ActionMap{
		"assignee":  action.ActionAssignableUsers(pr_createCmd).UniqueList(","),
		"base":      action.ActionBranches(pr_createCmd),
		"body":      action.ActionBody(pr_createCmd),
		"body-file": carapace.ActionFiles(),
		"head":      action.ActionBranches(pr_createCmd),
		"label":     action.ActionLabels(pr_createCmd).UniqueList(","),
		"milestone": action.ActionMilestones(pr_createCmd),
		"project":   action.ActionProjects(pr_createCmd, action.ProjectOpts{Open: true}),
		"reviewer":  action.ActionAssignableUsers(pr_createCmd).UniqueList(","),
		"template":  action.ActionPullRequestTemplates(pr_createCmd),
	})
}
