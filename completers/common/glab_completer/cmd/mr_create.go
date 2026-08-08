package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a new merge request.",
	Aliases: []string{"new"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_createCmd).Standalone()

	mr_createCmd.Flags().Bool("allow-collaboration", false, "Allow commits from other members. Set to true/false to override project defaults, or omit to use project settings.")
	mr_createCmd.Flags().StringSliceP("assignee", "a", nil, "Assign merge request to people by their `usernames`. Multiple usernames can be comma-separated or specified by repeating the flag.")
	mr_createCmd.Flags().Bool("auto-merge", false, "Set the merge request to merge when all merge checks pass.")
	mr_createCmd.Flags().Bool("copy-issue-labels", false, "Copy labels from issue to the merge request. Used with --related-issue.")
	mr_createCmd.Flags().Bool("create-source-branch", false, "Create a source branch if it does not exist.")
	mr_createCmd.Flags().StringP("description", "d", "", "Supply a description for the merge request. Set to \"-\" to open an editor.")
	mr_createCmd.Flags().Bool("draft", false, "Mark merge request as a draft.")
	mr_createCmd.Flags().BoolP("fill", "f", false, "Do not prompt for title or description, and just use commit info. Sets `push` to `true`, and pushes the branch.")
	mr_createCmd.Flags().Bool("fill-commit-body", false, "Fill description with each commit body when multiple commits. Can only be used with --fill.")
	mr_createCmd.Flags().StringP("head", "H", "", "Select another head repository using the `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format, the project ID, or the full URL.")
	mr_createCmd.Flags().StringSliceP("label", "l", nil, "Add label by name. Multiple labels can be comma-separated or specified by repeating the flag.")
	mr_createCmd.Flags().StringP("milestone", "m", "", "The global ID or title of a milestone to assign.")
	mr_createCmd.Flags().Bool("no-editor", false, "Don't open editor to enter a description. If true, uses prompt. Defaults to false.")
	mr_createCmd.Flags().Bool("push", false, "Push committed changes after creating merge request. Make sure you have committed changes.")
	mr_createCmd.Flags().Bool("recover", false, "Save the options to a file if the merge request creation fails. If the file exists, the options are loaded from the recovery file. (EXPERIMENTAL)")
	mr_createCmd.Flags().StringP("related-issue", "i", "", "Create a merge request for an issue. If --title is not provided, uses the issue title.")
	mr_createCmd.Flags().Bool("remove-source-branch", false, "Remove source branch on merge. Set to true/false to override project defaults, or omit to use project settings.")
	mr_createCmd.Flags().StringSlice("reviewer", nil, "Request review from users by their `usernames`. Multiple usernames can be comma-separated or specified by repeating the flag.")
	mr_createCmd.Flags().Bool("signoff", false, "Append a DCO signoff to the merge request description.")
	mr_createCmd.Flags().StringP("source-branch", "s", "", "Create a merge request from this branch. Default is the current branch.")
	mr_createCmd.Flags().Bool("squash-before-merge", false, "Squash commits into a single commit when merging. Set to true/false to override project defaults, or omit to use project settings.")
	mr_createCmd.Flags().StringP("target-branch", "b", "", "The target or base branch into which you want your code merged into.")
	mr_createCmd.Flags().String("target-project", "", "Add target project by id, OWNER/REPO, or GROUP/NAMESPACE/REPO.")
	mr_createCmd.Flags().String("template", "", "Name of a template in '.gitlab/merge_request_templates/' to pre-populate the description. The '.md' extension is optional. Templates are loaded from the local repository only.")
	mr_createCmd.Flags().StringP("title", "t", "", "Supply a title for the merge request.")
	mr_createCmd.Flags().BoolP("web", "w", false, "Continue merge request creation in a browser.")
	mr_createCmd.Flags().Bool("wip", false, "Mark merge request as a draft. Alternative to --draft.")
	mr_createCmd.Flags().BoolP("yes", "y", false, "Skip submission confirmation prompt. Use --fill to skip all optional prompts.")
	mr_createCmd.Flag("target-project").Hidden = true
	mr_createCmd.Flag("target-project").Hidden = true
	mrCmd.AddCommand(mr_createCmd)

	// TODO target-project completion
	carapace.Gen(mr_createCmd).FlagCompletion(carapace.ActionMap{
		"assignee":      action.ActionProjectMembers(mr_createCmd).UniqueList(","),
		"label":         action.ActionLabels(mr_createCmd).UniqueList(","),
		"milestone":     action.ActionMilestones(mr_createCmd),
		"related-issue": action.ActionIssues(mr_createCmd, "opened"),
		"reviewer":      action.ActionProjectMembers(mr_createCmd).UniqueList(","),
		"source-branch": action.ActionBranches(mr_createCmd),
		"target-branch": action.ActionBranches(mr_createCmd),
	})
}
